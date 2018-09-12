# Chapter 12: Reflection

<!-- TOC -->

- [12.1. Why Reflection?](#121-why-reflection)
- [12.2. reflect.Type and reflect.Value](#122-reflecttype-and-reflectvalue)
- [12.3. Display, a Recursive Value Printer](#123-display-a-recursive-value-printer)
- [12.4. Example: Encoding S-Expressions](#124-example-encoding-s-expressions)
- [12.5. Setting Variables with reflect.Value](#125-setting-variables-with-reflectvalue)
- [12.6. Example: Decoding S-Expressions](#126-example-decoding-s-expressions)
- [12.7. Accessing Struct Field Tags](#127-accessing-struct-field-tags)
- [12.8. Displaying the Methods of a Type](#128-displaying-the-methods-of-a-type)
- [12.9. A Word of Caution](#129-a-word-of-caution)

<!-- /TOC -->


Go provides a mechanism to update variables and inspect their values at run time, to call their methods, and to apply the operations intrinsic to their representation, all without knowing their types at compile time. This mechanism is called *reflection*. Reflection also lets us treat types themselves as first-class values.

In this chapter, we'll explore Go's reflection features to see how they increase the expressiveness of the language, and in particular how they are crucial to the implementation of two important APIs: string formatting provided by `fmt`, and protocol encoding provided by packages like `encoding/json` and `encoding/xml`. Reflection is also essential to the template mechanism provided by the `text/template` and `html/template` packages we saw in Section 4.6. However, reflection is complex to reason about and not for casual use, so although these packages are implemented using reflection, they do not expose reflection in their own APIs.


## 12.1. Why Reflection? 

Sometimes we need to write a function capable of dealing uniformly with values of types that don't satisfy a common interface, don't have a known representation, or don't exist at the time we design the function; or even all three.

A familiar example is the formatting logic within `fmt.Fprintf`, which can usefully print an arbitrary value of any type, even a user-defined one. Let's try to implement a function like it using what we know already. For simplicity, our function will accept one argument and will return the result as a string like `fmt.Sprint` does, so we'll call it `Sprint`.

We start with a type switch that tests whether the argument defines a `String` method, and call it if so. We then add switch cases that test the value's dynamic type against each of the basic types (`string`, `int`, `bool`, and so on) and perform the appropriate formatting operation in each case.
```go
    func Sprint(x interface{}) string {
        type stringer interface {
            String() string
        }
        switch x := x.(type) {
        case stringer:
            return x.String()
        case string:
            return x
        case int:
            return strconv.Itoa(x)
        // ...similar cases for int16, uint32, and so on...
        case bool:
            if x {
                return "true"
            }
            return "false"
        default:
            // array, chan, func, map, pointer, slice, struct
            return "???"
        } 
    }
```
But how do we deal with other types, like `[]float64`, `map[string][]string`, and so on? We could add more cases, but the number of such types is infinite. And what about named types, like `url.Values`? Even if the type switch had a case for its underlying type `map[string][]string`, it wouldn't match `url.Values` because the two types are not identical, and the type switch cannot include a case for each type like `url.Values` because that would require this library to depend upon its clients.

Without a way to inspect the representation of values of unknown types, we quickly get stuck. What we need is reflection.


## 12.2. `reflect.Type` and `reflect.Value` 

Reflection is provided by the `reflect` package. It defines two important types, `Type` and `Value`. A `Type` represents a Go type. It is an interface with many methods for discriminating among types and inspecting their components, like the fields of a struct or the parameters of a function. The sole implementation of `reflect.Type` is the type descriptor (§7.5), the same entity that identifies the dynamic type of an interface value.

The `reflect.TypeOf` function accepts any `interface{}` and returns its dynamic type as a `reflect.Type`:
```go
    t := reflect.TypeOf(3)  // a reflect.Type
    fmt.Println(t.String()) // "int"
    fmt.Println(t)          // "int"
```
The `TypeOf(3)` call above assigns the value `3` to the `interface{}` parameter. Recall from Section 7.5 that an assignment from a concrete value to an interface type performs an implicit interface conversion, which creates an interface value consisting of two components: its *dynamic type* is the operand's type (`int`) and its *dynamic value* is the operand's value (`3`).

Because `reflect.TypeOf` returns an interface value's dynamic type, it always returns a concrete type. So, for example, the code below prints `"*os.File"`, not `"io.Writer"`. Later, we will see that `reflect.Type` is capable of representing interface types too.
```go
    var w io.Writer = os.Stdout
    fmt.Println(reflect.TypeOf(w)) // "*os.File"
```
Notice that `reflect.Type` satisfies `fmt.Stringer`. Because printing the dynamic type of an interface value is useful for debugging and logging, `fmt.Printf` provides a shorthand, `%T`, that uses `reflect.TypeOf` internally:
```go
    fmt.Printf("%T\n", 3) // "int"
```
The other important type in the `reflect` package is `Value`. A `reflect.Value` can hold a value of any type. The `reflect.ValueOf` function accepts any `interface{}` and returns a `reflect.Value` containing the interface's dynamic value. As with `reflect.TypeOf`, the results of `reflect.ValueOf` are always concrete, but a `reflect.Value` can hold interface values too.
```go
    v := reflect.ValueOf(3) // a reflect.Value
    fmt.Println(v)          // "3"
    fmt.Printf("%v\n", v)   // "3"
    fmt.Println(v.String()) // NOTE: "<int Value>"
```
Like `reflect.Type`, `reflect.Value` also satisfies `fmt.Stringer`, but unless the `Value` holds a string, the result of the `String` method reveals only the type. Instead, use the `fmt` package's `%v` verb, which treats `reflect.Values` specially.

Calling the `Type` method on a `Value` returns its type as a `reflect.Type`:
```go
    t := v.Type()           // a reflect.Type
    fmt.Println(t.String()) // "int"
```
The inverse operation to `reflect.ValueOf` is the `reflect.Value.Interface` method. It returns an `interface{}` holding the same concrete value as the `reflect.Value`:
```go
  v := reflect.ValueOf(3) // a reflect.Value
  x := v.Interface()      // an interface{}
  i := x.(int)            // an int
  fmt.Printf("%d\n", i)   // "3"
```
A `reflect.Value` and an `interface{}` can both hold arbitrary values. The difference is that an empty interface hides the representation and intrinsic operations of the value it holds and exposes none of its methods, so unless we know its dynamic type and use a type assertion to peer inside it (as we did above), there is little we can do to the value within. In contrast, a `Value` has many methods for inspecting its contents, regardless of its type. Let's use them for our second attempt at a general formatting function, which we'll call `format.Any`.

Instead of a type switch, we use `reflect.Value`'s Kind method to discriminate the cases. Although there are infinitely many types, there are only a finite number of kinds of type: the basic types `Bool`, `String`, and all the numbers; the aggregate types `Array` and `Struct`; the reference types `Chan`, `Func`, `Ptr`, `Slice`, and `Map`; `Interface` types; and finally `Invalid`, meaning no value at all. (The zero value of a `reflect.Value` has kind `Invalid`.)
```go
// gopl.io/ch12/format
// Package format provides an Any function that can format any value.
package format

import (
	"reflect"
	"strconv"
)

// Any formats any value as a string.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
```
So far, our function treats each value as an indivisible thing with no internal structure; hence `formatAtom`. For aggregate types (structs and arrays) and interfaces it prints only the *type* of the value, and for reference types (channels, functions, pointers, slices, and maps), it prints the type and the reference address in hexadecimal. This is less than ideal but still a major improvement, and since `Kind` is concerned only with the underlying representation, `format.Any` works for named types too. For example:
```go
    var x int64 = 1
    var d time.Duration = 1 * time.Nanosecond

    fmt.Println(format.Any(x))                  // "1"
    fmt.Println(format.Any(d))                  // "1"
    fmt.Println(format.Any([]int64{x}))         // "[]int64 0x8202b87b0"
    fmt.Println(format.Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"
```

## 12.3. `Display`, a Recursive Value Printer 

Next we'll take a look at how to improve the display of composite types. Rather than try to copy `fmt.Sprint` exactly, we'll build a debugging utility function called `Display` that, given an arbitrarily complex value `x`, prints the complete structure of that value, labeling each element with the path by which it was found. Let's start with an example.
```go
    e, _ := eval.Parse("sqrt(A / pi)")
    Display("e", e)
```
In the call above, the argument to Display is a syntax tree from the expression evaluator in Section 7.9. The output of `Display` is shown below:
```go
    Display e (eval.call):
    e.fn = "sqrt"
    e.args[0].type = eval.binary
    e.args[0].value.op = 47
    e.args[0].value.x.type = eval.Var
    e.args[0].value.x.value = "A"
    e.args[0].value.y.type = eval.Var
    e.args[0].value.y.value = "pi"
```
Where possible, you should avoid exposing reflection in the API of a package. We'll define an unexported function `display` to do the real work of the recursion, and export `Display`, a simple wrapper around it that accepts an `interface{}` parameter:
```go
// gopl.io/ch12/display
    func Display(name string, x interface{}) {
      fmt.Printf("Display %s (%T):\n", name, x)
      display(name, reflect.ValueOf(x))
    }
```
In `display`, we'll use the `formatAtom` function we defined earlier to print elementary values (basic types, functions, and channels) but we'll use the methods of `reflect.Value` to recursively display each component of a more complex type. As the recursion descends, the `path` string, which initially describes the starting value (for instance, `"e"`), will be augmented to indicate how we reached the current value (for instance, `"e.args[0].value"`).

Since we're no longer pretending to implement `fmt.Sprint`, we will use the `fmt` package to keep our example short.
```go
// gopl.io/ch12/display
func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
```
Let's discuss the cases in order:

- *Slices and Arrays*: 
  - The logic is the same for both. The `Len` method returns the number of elements of a slice or array value, and `Index(i)` retrieves the element at index `i`, also as a `reflect.Value`; it panics if `i` is out of bounds. These are analogous to the built-in `len(a)` and `a[i]` operations on sequences. The `display` function recursively invokes itself on each element of the sequence, appending the subscript notation `"[i]"` to the path.
  - Although `reflect.Value` has many methods, only a few are safe to call on any given value. For example, the `Index` method may be called on values of kind `Slice`, `Array`, or `String`, but panics for any other kind.
- *Structs*: 
  - The `NumField` method reports the number of fields in the struct, and `Field(i)` returns the value of the *i*-th field as a `reflect.Value`. The list of fields includes ones promoted from anonymous fields. To append the field selector notation `".f"` to the path, we must obtain the `reflect.Type` of the struct and access the name of its *i*-th field.
- *Maps*: 
  - The `MapKeys` method returns a slice of `reflect.Values`, one per map key. As usual when iterating over a map, the order is undefined. `MapIndex(key)` returns the value corresponding to key. We append the subscript notation `"[key]"` to the path. (We're cutting a corner here. The type of a map key isn't restricted to the types formatAtom handles best; arrays, structs, and interfaces can also be valid map keys. Extending this case to print the key in full is Exercise 12.1.)
- *Pointers*: 
  - The `Elem` method returns the variable pointed to by a pointer, again as a `reflect.Value.` This operation would be safe even if the pointer value is `nil,` in which case the result would have kind `Invalid,` but we use `IsNil` to detect nil pointers explicitly so we can print a more appropriate message. We prefix the path with a `"*"` and parenthesize it to avoid ambiguity.
- *Interfaces*: 
  - Again, we use `IsNil` to test whether the interface is nil, and if not, we retrieve its dynamic value using `v.Elem()` and print its type and value.

Now that our `Display` function is complete, let's put it to work. The `Movie` type below is a slight variation on the one in Section 4.5:
```go
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
```
Let's declare a value of this type and see what `Display` does with it:
```go
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
```
The call `Display("strangelove", strangelove)` prints:
```
    Display strangelove (display.Movie):
    strangelove.Title = "Dr. Strangelove"
    strangelove.Subtitle = "How I Learned to Stop Worrying and Love the Bomb"
    strangelove.Year = 1964
    strangelove.Color = false
    strangelove.Actor["Gen. Buck Turgidson"] = "George C. Scott"
    strangelove.Actor["Brig. Gen. Jack D. Ripper"] = "Sterling Hayden"
    strangelove.Actor["Maj. T.J. \"King\" Kong"] = "Slim Pickens"
    strangelove.Actor["Dr. Strangelove"] = "Peter Sellers"
    strangelove.Actor["Grp. Capt. Lionel Mandrake"] = "Peter Sellers"
    strangelove.Actor["Pres. Merkin Muffley"] = "Peter Sellers"
    strangelove.Oscars[0] = "Best Actor (Nomin.)"
    strangelove.Oscars[1] = "Best Adapted Screenplay (Nomin.)"
    strangelove.Oscars[2] = "Best Director (Nomin.)"
    strangelove.Oscars[3] = "Best Picture (Nomin.)"
    strangelove.Sequel = nil
```
We can use `Display` to display the internals of library types, such as `*os.File`:
```go
    Display("os.Stderr", os.Stderr)
    // Output:
    // Display os.Stderr (*os.File):
    // (*(*os.Stderr).file).fd = 2
    // (*(*os.Stderr).file).name = "/dev/stderr"
    // (*(*os.Stderr).file).nepipe = 0
```
Notice that even unexported fields are visible to reflection. Beware that the particular output of this example may vary across platforms and may change over time as libraries evolve. (Those fields are private for a reason!) We can even apply `Display` to a `reflect.Value` and watch it traverse the internal representation of the type descriptor for `*os.File`. The output of the call `Display("rV", reflect.ValueOf(os.Stderr))` is shown below, though of course your mileage may vary:
```
    Display rV (reflect.Value):
    (*rV.typ).size = 8
    (*rV.typ).hash = 871609668
    (*rV.typ).align = 8
    (*rV.typ).fieldAlign = 8
    (*rV.typ).kind = 22
    (*(*rV.typ).string) = "*os.File"
    (*(*(*rV.typ).uncommonType).methods[0].name) = "Chdir" 
    (*(*(*(*rV.typ).uncommonType).methods[0].mtyp).string) = "func() error" 
    (*(*(*(*rV.typ).uncommonType).methods[0].typ).string) = "func(*os.File) error" 
    ...
```
Observe the difference between these two examples:
```go
    var i interface{} = 3

    Display("i", i)
    // Output:
    // Display i (int):
    // i = 3

    Display("&i", &i)
    // Output:
    // Display &i (*interface {}):
    // (*&i).type = int
    // (*&i).value = 3
```
In the first example, `Display` calls `reflect.ValueOf(i)`, which returns a value of kind `Int`. As we mentioned in Section 12.2, `reflect.ValueOf` always returns a Value of a concrete type since it extracts the contents of an interface value.

In the second example, Display calls `reflect.ValueOf(&i)`, which returns a pointer to `i`, of `kindPtr`. The switch case for `Ptr` calls `Elem` on this value, which returns a `Value` representing the *variable* `i` itself, of kind `Interface`. A `Value` obtained indirectly, like this one, may represent any value at all, including interfaces. The `display` function calls itself recursively and this time, it prints separate components for the interface's dynamic type and value.

As currently implemented, `Display` will never terminate if it encounters a cycle in the object graph, such as this linked list that eats its own tail:
```go
    // a struct that points to itself
    type Cycle struct{ Value int; Tail *Cycle }
    var c Cycle
    c = Cycle{42, &c}
    Display("c", c)
```
Display prints this ever-growing expansion:
```go
    Display c (display.Cycle):
    c.Value = 42
    (*c.Tail).Value = 42 
    (*(*c.Tail).Tail).Value = 42 
    (*(*(*c.Tail).Tail).Tail).Value = 42 
    ...ad infinitum...
```
Many Go programs contain at least some cyclic data. Making Display robust against such cycles is tricky, requiring additional bookkeeping to record the set of references that have been followed so far; it is costly too. A general solution requires unsafe language features, as we will see in Section 13.3.

Cycles pose less of a problem for `fmt.Sprint` because it rarely tries to print the complete structure. For example, when it encounters a pointer, it breaks the recursion by printing the pointer's numeric value. It can get stuck trying to print a slice or map that contains itself as an element, but such rare cases do not warrant the considerable extra trouble of handling cycles.

### Exercises
- **Exercise 12.1**: Extend `Display` so that it can display maps whose keys are structs or arrays.
- **Exercise 12.2**: Make `display` safe to use on cyclic data structures by bounding the number of steps it takes before abandoning the recursion. (In Section 13.3, we'll see another way to detect cycles.)


## 12.4. Example: Encoding S-Expressions 

`Display` is a debugging routine for displaying structured data, but it's not far short of being able to encode or *marshal* arbitrary Go objects as messages in a portable notation suitable for inter-process communication.

As we saw in Section 4.5, Go's standard library supports a variety of formats, including JSON, XML, and ASN.1. Another notation that is still widely used is *S-expressions*, the syntax of Lisp. Unlike the other notations, S-expressions are not supported by the Go standard library, not least because they have no universally accepted definition, despite several attempts at standardization and the existence of many implementations.

In this section, we'll define a package that encodes arbitrary Go objects using an S-expression notation that supports the following constructs:
```
42            integer
"hello"       string (with Go-style quotation)
foo           symbol (an unquoted name)
(1 2 3)       list   (zero or more items enclosed in parentheses)
```
Booleans are traditionally encoded using the symbol `t` for true, and the empty list `()` or the symbol `nil` for false, but for simplicity, our implementation ignores them. It also ignores channels and functions, since their state is opaque to reflection. And it ignores real and complex floating-point numbers and interfaces. Adding support for them is Exercise 12.3.

We'll encode the types of Go using S-expressions as follows. Integers and strings are encoded in the obvious way. Nil values are encoded as the symbol `nil`. Arrays and slices are encoded using list notation.

Structs are encoded as a list of field bindings, each field binding being a two-element list whose first element (a symbol) is the field name and whose second element is the field value. Maps too are encoded as a list of pairs, with each pair being the key and value of one map entry. Traditionally, S-expressions represent lists of key/value pairs using a single *cons* cell `(key . value)` for each pair, rather than a two-element list, but to simplify the decoding we'll ignore dotted list notation.

Encoding is done by a single recursive function, `encode`, shown below. Its structure is essentially the same as that of `Display` in the previous section:
```go
// gopl.io/ch12/sexpr
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
```
The `Marshal` function wraps the encoder in an API similar to those of the other `encoding/...` packages:
```go
// gopl.io/ch12/sexpr/encode
    // Marshal encodes a Go value in S-expression form.
    func Marshal(v interface{}) ([]byte, error) {
      var buf bytes.Buffer
      if err := encode(&buf, reflect.ValueOf(v)); err != nil {
        return nil, err
      }
      return buf.Bytes(), nil
    }
```
Here's the output of `Marshal` applied to the `strangelove` variable from Section 12.3:
```
    ((Title "Dr. Strangelove") (Subtitle "How I Learned to Stop Worrying and Lo
    ve the Bomb") (Year 1964) (Actor (("Grp. Capt. Lionel Mandrake" "Peter Sell
    ers") ("Pres. Merkin Muffley" "Peter Sellers") ("Gen. Buck Turgidson" "Geor
    ge C. Scott") ("Brig. Gen. Jack D. Ripper" "Sterling Hayden") ("Maj. T.J. \
    "King\" Kong" "Slim Pickens") ("Dr. Strangelove" "Peter Sellers"))) (Oscars
    ("Best Actor (Nomin.)" "Best Adapted Screenplay (Nomin.)" "Best Director (N
    omin.)" "Best Picture (Nomin.)")) (Sequel nil))
```
The whole output appears on one long line with minimal spaces, making it hard to read. Here's the same output manually formatted according to S-expression conventions. Writing a pretty-printer for S-expressions is left as a (challenging) exercise; the download from `gopl.io` includes a simple version.
```
    ((Title "Dr. Strangelove")
    (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
    (Year 1964)
    (Actor (("Grp. Capt. Lionel Mandrake" "Peter Sellers")
            ("Pres. Merkin Muffley" "Peter Sellers")
            ("Gen. Buck Turgidson" "George C. Scott")
            ("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
            ("Maj. T.J. \"King\" Kong" "Slim Pickens")
            ("Dr. Strangelove" "Peter Sellers")))
    (Oscars ("Best Actor (Nomin.)"
              "Best Adapted Screenplay (Nomin.)"
              "Best Director (Nomin.)"
              "Best Picture (Nomin.)"))
    (Sequel nil))
```
Like the f`mt.Print`, `json.Marshal`, and `Display` functions, `sexpr.Marshal` will loop forever if called with cyclic data.

In Section 12.6, we'll sketch out the implementation of the corresponding S-expression decoding function, but before we get there, we'll first need to understand how reflection can be used to update program variables.

### Exercises
- **Exercise 12.3**: Implement the missing cases of the encode function. Encode booleans as t and nil, floating-point numbers using Go's notation, and complex numbers like `1+2i` as `#C(1.02.0)`. Interfaces can be encoded as a pair of a type name and a value, for instance `("[]int"(123))`, but beware that this notation is ambiguous: the `reflect.Type.String` method may return the same string for different types.
- **Exercise 12.4**: Modify encode to pretty-print the S-expression in the style shown above. 
- **Exercise 12.5**: Adapt encode to emit JSON instead of S-expressions. Test your encoder using the standard decoder, `json.Unmarshal`.
- **Exercise 12.6**: Adapt encode so that, as an optimization, it does not encode a field whose value is the zero value of its type.
- **Exercise 12.7**: Create a streaming API for the S-expression decoder, following the style of `json.Decoder` (§4.5).


## 12.5. Setting Variables with `reflect.Value` 

So far, reflection has only *interpreted* values in our program in various ways. The point of this section, however, is to *change* them.

Recall that some Go expressions like `x`, `x.f[1]`, and `*p` denote variables, but others like `x + 1` and `f(2)` do not. A variable is an *addressable* storage location that contains a value, and its value may be updated through that address.

A similar distinction applies to `reflect.Values`. Some are addressable; others are not. Consider the following declarations:
```go
x := 2                   // value    type   variable?
a := reflect.ValueOf(2)  // 2        int    no
b := reflect.ValueOf(x)  // 2        int    no
c := reflect.ValueOf(&x) // &x       *int   no
d := c.Elem()            // 2        int yes (x)
```
The value within a is not addressable. It is merely a copy of the integer 2. The same is true of `b`. The value within `c` is also non-addressable, being a copy of the pointer value `&x`. In fact, no `reflect.Value` returned by `reflect.ValueOf(x)` is addressable. But `d`, derived from `c` by dereferencing the pointer within it, refers to a variable and is thus addressable. We can use this approach, calling `reflect.ValueOf(&x).Elem()`, to obtain an addressable `Value` for any variable `x`.

We can ask a `reflect.Value` whether it is addressable through its `CanAddr` method:
```go
    fmt.Println(a.CanAddr()) // "false"
    fmt.Println(b.CanAddr()) // "false"
    fmt.Println(c.CanAddr()) // "false"
    fmt.Println(d.CanAddr()) // "true"
```
We obtain an addressable `reflect.Value` whenever we indirect through a pointer, even if we started from a non-addressable Value. All the usual rules for addressability have analogs for reflection. For example, since the slice indexing expression `e[i]` implicitly follows a pointer, it is addressable even if the expression `e` is not. By analogy, `reflect.ValueOf(e).Index(i)` refers to a variable, and is thus addressable even if `reflect.ValueOf(e)` is not.

To recover the variable from an addressable `reflect.Value` requires three steps. First, we call `Addr()`, which returns a `Value` holding a pointer to the variable. Next, we call `Interface()` on this Value, which returns an `interface{}` value containing the pointer. Finally, if we know the type of the variable, we can use a type assertion to retrieve the contents of the interface as an ordinary pointer. We can then update the variable through the pointer:
```go
x := 2
d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
px := d.Addr().Interface().(*int) // px := &x
*px = 3                           // x = 3
fmt.Println(x)                    // "3"
```
Or, we can update the variable referred to by an addressable `reflect.Value` directly, without using a pointer, by calling the `reflect.Value.Set` method:
```go
    d.Set(reflect.ValueOf(4))
    fmt.Println(x) // "4"
```
The same checks for assignability that are ordinarily performed by the compiler are done at run time by the `Set` methods. Above, the variable and the value both have type `int`, but if the variable had been an `int64`, the program would panic, so it's crucial to make sure the value is assignable to the type of the variable:
```go
    d.Set(reflect.ValueOf(int64(5))) // panic: int64 is not assignable to int
```
And of course calling `Set` on a non-addressable `reflect.Value` panics too:
```go
    x := 2
    b := reflect.ValueOf(x)
    b.Set(reflect.ValueOf(3)) // panic: Set using unaddressable value
```
There are variants of `Set` specialized for certain groups of basic types: `SetInt`, `SetUint`, `SetString`, `SetFloat`, and so on:
```go
    d := reflect.ValueOf(&x).Elem()
    d.SetInt(3)
    fmt.Println(x) // "3"
```
In some ways these methods are more forgiving. `SetInt`, for example, will succeed so long as the variable's type is some kind of signed integer, or even a named type whose underlying type is a signed integer, and if the value is too large it will be quietly truncated to fit. But tread carefully: calling `SetInt` on a `reflect.Value` that refers to an `interface{}` variable will panic, even though `Set` would succeed.
```go
x := 1
rx := reflect.ValueOf(&x).Elem()
rx.SetInt(2)                     // OK, x = 2
rx.Set(reflect.ValueOf(3))       // OK, x = 3
rx.SetString("hello")            // panic: string is not assignable to int
rx.Set(reflect.ValueOf("hello")) // panic: string is not assignable to int

var y interface{}
ry := reflect.ValueOf(&y).Elem()
ry.SetInt(2)                     // panic: SetInt called on interface Value
ry.Set(reflect.ValueOf(3))       // OK, y = int(3)
ry.SetString("hello")            // panic: SetString called on interface Value
ry.Set(reflect.ValueOf("hello")) // OK, y = "hello"
```
When we applied `Display` to `os.Stdout`, we found that reflection can read the values of unexported struct fields that are inaccessible according to the usual rules of the language, like the `fd` `int` field of an `os.File` struct on a Unix-like platform. However, reflection cannot update such values:
```go
    stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
    fmt.Println(stdout.Type())                  // "os.File"
    fd := stdout.FieldByName("fd")
    fmt.Println(fd.Int()) // "1"
    fd.SetInt(2)          // panic: unexported field
```
An addressable `reflect.Value` records whether it was obtained by traversing an unexported struct field and, if so, disallows modification. Consequently, `CanAddr` is not usually the right check to use before setting a variable. The related method `CanSet` reports whether a `reflect.Value` is addressable *and* settable:
```go
    fmt.Println(fd.CanAddr(), fd.CanSet()) // "true false"
```


## 12.6. Example: Decoding S-Expressions 

For each `Marshal` function provided by the standard library's `encoding/...` packages, there is a corresponding `Unmarshal` function that does decoding. For example, as we saw in Section 4.5, given a byte slice containing JSON-encoded data for our `Movie` type (§12.3), we can decode it like this:
```go
    data := []byte{/* ... */}
    var movie Movie
    err := json.Unmarshal(data, &movie)
```
The `Unmarshal` function uses reflection to modify the fields of the existing `movie` variable, creating new maps, structs, and slices as determined by the type `Movie` and the content of the incoming data.

Let's now implement a simple Unmarshal function for S-expressions, analogous to the standard `json.Unmarshal` function used above, and the inverse of our earlier `sexpr.Marshal`. We must caution you that a robust and general implementation requires substantially more code than will comfortably fit in this example, which is already long, so we have taken many shortcuts. We support only a limited subset of S-expressions and do not handle errors gracefully. The code is intended to illustrate reflection, not parsing.

The lexer uses the `Scanner` type from the `text/scanner` package to break an input stream into a sequence of tokens such as comments, identifiers, string literals, and numeric literals. The scanner's `Scan` method advances the scanner and returns the kind of the next token, which has type rune. Most tokens, like `'('`, consist of a single rune, but the `text/scanner` package represents the kinds of the multi-character tokens `Ident`, `String`, and `Int` using small negative values of type `rune`. Following a call to `Scan` that returns one of these kinds of token, the scanner's `TokenText` method returns the text of the token.

Since a typical parser may need to inspect the current token several times, but the `Scan` method advances the scanner, we wrap the scanner in a helper type called `lexer` that keeps track of the token most recently returned by `Scan`.
```go
// gopl.io/ch12/sexpr
type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}
```
Now let's turn to the parser. It consists of two principal functions. The first of these, `read`, reads the S-expression that starts with the current token and updates the variable referred to by the addressable `reflect.Value` `v`.
```go
func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		// The only valid identifiers are
		// "nil" and struct field names.
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		v.SetInt(int64(i))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next() // consume ')'
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}
```
Our S-expressions use identifiers for two distinct purposes, struct field names and the `nil` value for a pointer. The read function only handles the latter case. When it encounters the `scanner.Ident` `"nil"`, it sets `v` to the zero value of its type using the `reflect.Zero` function. For any other identifier, it reports an error. The `readList` function, which we'll see in a moment, handles identifiers used as struct field names.

A `'('` token indicates the start of a list. The second function, `readList`, decodes a list into a variable of composite type (a map, struct, slice, or array) depending on what kind of Go variable we're currently populating. In each case, the loop keeps parsing items until it encounters the matching close parenthesis, `')'`, as detected by the endList function.

The interesting part is the recursion. The simplest case is an array. Until the closing `')'` is seen, we use `Index` to obtain the variable for each array element and make a recursive call to `read` to populate it. As in many other error cases, if the input data causes the decoder to index beyond the end of the array, the decoder panics. A similar approach is used for slices, except we must create a new variable for each element, populate it, then append it to the slice.

The loops for structs and maps must parse a `(key value)` sublist on each iteration. For structs, the key is a symbol identifying the field. Analogous to the case for arrays, we obtain the existing variable for the struct field using `FieldByName` and make a recursive call to populate it. For maps, the key may be of any type, and analogous to the case for slices, we create a new variable, recursively populate it, and finally insert the new key/value pair into the map.
```go
func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
```
Finally, we wrap up the parser in an exported function `Unmarshal`, shown below, that hides some of the rough edges of the implementation. Errors encountered during parsing result in a panic, so `Unmarshal` uses a deferred call to recover from the panic (§5.10) and return an error message instead.
```go
func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next() // get the first token
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}
```
A production-quality implementation should never panic for any input and should report an informative error for every mishap, perhaps with a line number or offset. Nonetheless, we hope this example conveys some idea of what's happening under the hood of the packages like `encoding/json`, and how you can use reflection to populate data structures.

### Exercises
- **Exercise 12.8**: The `sexpr.Unmarshal` function, like `json.Marshal`, requires the complete input in a byte slice before it can begin decoding. Define a `sexpr.Decoder` type that, like `json.Decoder`, allows a sequence of values to be decoded from an `io.Reader`. Change `sexpr.Unmarshal` to use this new type.
- **Exercise 12.9**: Write a token-based API for decoding S-expressions, following the style of `xml.Decoder` (§7.14). You will need five types of tokens: `Symbol`, `String`, `Int`, `StartList`, and `EndList`.
- **Exercise 12.10**: Extend `sexpr.Unmarshal` to handle the booleans, floating-point numbers, and interfaces encoded by your solution to Exercise 12.3. (Hint: to decode interfaces, you will need a mapping from the name of each supported type to its `reflect.Type`.)


## 12.7. Accessing Struct Field Tags 

In Section 4.5 we used struct *field tags* to modify the JSON encoding of Go struct values. The json field tag lets us choose alternative field names and suppress the output of empty fields. In this section, we'll see how to access field tags using reflection.

In a web server, the first thing most HTTP handler functions do is extract the request parameters into local variables. We'll define a utility function, `params.Unpack`, that uses struct field tags to make writing HTTP handlers (§7.7) more convenient.

First, we'll show how it's used. The `search` function below is an HTTP handler. It defines a variable called `data` of an anonymous struct type whose fields correspond to the HTTP request parameters. The struct's field tags specify the parameter names, which are often short and cryptic since space is precious in a URL. The `Unpack` function populates the struct from the request so that the parameters can be accessed conveniently and with an appropriate type.
```go
// gopl.io/ch12/params
import "gopl.io/ch12/params"

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}
```
The `Unpack` function below does three things. First, it calls `req.ParseForm()` to parse the request. Thereafter, `req.Form` contains all the parameters, regardless of whether the HTTP client used the GET or the POST request method.

Next, `Unpack` builds a mapping from the *effective* name of each field to the variable for that field. The effective name may differ from the actual name if the field has a tag. The `Field` method of `reflect.Type` returns a `reflect.StructField` that provides information about the type of each field such as its name, type, and optional tag. The Tag field is a `reflect.StructTag`, which is a string type that provides a `Get` method to parse and extract the substring for a particular key, such as `http:"..."` in this case.
```go
// gopl.io/ch12/params
// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}
```
Finally, `Unpack` iterates over the name/value pairs of the HTTP parameters and updates the corresponding struct fields. Recall that the same parameter name may appear more than once. If this happens, and the field is a slice, then all the values of that parameter are accumulated into the slice. Otherwise, the field is repeatedly overwritten so that only the last value has any effect.

The populate function takes care of setting a single field `v` (or a single element of a slice field) from a parameter value. For now, it supports only strings, signed integers, and booleans. Supporting other types is left as an exercise.
```go
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
```
If we add the `server` handler to a web server, this might be a typical session:
```
    $ go build gopl.io/ch12/search
    $ ./search &
    $ ./fetch 'http://localhost:12345/search'
    Search: {Labels:[] MaxResults:10 Exact:false}
    $ ./fetch 'http://localhost:12345/search?l=golang&l=programming'
    Search: {Labels:[golang programming] MaxResults:10 Exact:false}
    $ ./fetch 'http://localhost:12345/search?l=golang&l=programming&max=100'
    Search: {Labels:[golang programming] MaxResults:100 Exact:false}
    $ ./fetch 'http://localhost:12345/search?x=true&l=golang&l=programming'
    Search: {Labels:[golang programming] MaxResults:10 Exact:true}
    $ ./fetch 'http://localhost:12345/search?q=hello&x=123'
    x: strconv.ParseBool: parsing "123": invalid syntax
    $ ./fetch 'http://localhost:12345/search?q=hello&max=lots'
    max: strconv.ParseInt: parsing "lots": invalid syntax
```

### Exercises
- **Exercise 12.11**: Write the corresponding `Pack` function. Given a struct value, `Pack` should return a URL incorporating the parameter values from the struct.
- **Exercise 12.12**: Extend the field tag notation to express parameter validity requirements. For example, a string might need to be a valid email address or credit-card number, and an integer might need to be a valid US ZIP code. Modify `Unpack` to check these requirements.
- **Exercise 12.13**: Modify the S-expression encoder (§12.4) and decoder (§12.6) so that they honor the `sexpr:"..."` field tag in a similar manner to `encoding/json` (§4.5).


## 12.8. Displaying the Methods of a Type 

Our final example of reflection uses `reflect.Type` to print the type of an arbitrary value and enumerate its methods:
```go
// gopl.io/ch12/methods
// Print prints the method set of the value x.
func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}
```
Both `reflect.Type` and `reflect.Value` have a method called `Method`. Each `t.Method(i)` call returns an instance of `reflect.Method`, a struct type that describes the name and type of a single method. Each `v.Method(i)` call returns a `reflect.Value` representing a method value (§6.4), that is, a method bound to its receiver. Using the `reflect.Value.Call` method (which we don't have space to show here), it's possible to call `Values` of kind `Func` like this one, but this program needs only its Type.

Here are the methods belonging to two types, `time.Duration` and `*strings.Replacer`:
```go
    methods.Print(time.Hour)
    // Output:
    // type time.Duration
    // func (time.Duration) Hours() float64
    // func (time.Duration) Minutes() float64
    // func (time.Duration) Nanoseconds() int64
    // func (time.Duration) Seconds() float64
    // func (time.Duration) String() string

    methods.Print(new(strings.Replacer))
    // Output:
    // type *strings.Replacer
    // func (*strings.Replacer) Replace(string) string
    // func (*strings.Replacer) WriteString(io.Writer, string) (int, error)
```


## 12.9. A Word of Caution 

There is a lot more to the reflection API than we have space to show, but the preceding examples give an idea of what is possible. Reflection is a powerful and expressive tool, but it should be used with care, for three reasons.

The first reason is that reflection-based code can be fragile. For every mistake that would cause a compiler to report a type error, there is a corresponding way to misuse reflection, but whereas the compiler reports the mistake at build time, a reflection error is reported during execution as a panic, possibly long after the program was written or even long after it has started running.

If the `readList` function (§12.6), for example, should read a string from the input while populating a variable of type `int`, the call to `reflect.Value.SetString` will panic. Most programs that use reflection have similar hazards, and considerable care is required to keep track of the type, addressability, and settability of each `reflect.Value`.

The best way to avoid this fragility is to ensure that the use of reflection is fully encapsulated within your package and, if possible, avoid `reflect.Value` in favor of specific types in your package's API, to restrict inputs to legal values. If this is not possible, perform additional dynamic checks before each risky operation. As an example from the standard library, when `fmt.Printf` applies a verb to an inappropriate operand, it does not panic mysteriously but prints an informative error message. The program still has a bug, but it is easier to diagnose.
```go
    fmt.Printf("%d %s\n", "hello", 42) // "%!d(string=hello) %!s(int=42)"
```
Reflection also reduces the safety and accuracy of automated refactoring and analysis tools, because they can't determine or rely on type information.

The second reason to avoid reflection is that since types serve as a form of documentation and the operations of reflection cannot be subject to static type checking, heavily reflective code is often hard to understand. Always carefully document the expected types and other invariants of functions that accept an `interface{}` or a `reflect.Value`.

The third reason is that reflection-based functions may be one or two orders of magnitude slower than code specialized for a particular type. In a typical program, the majority of functions are not relevant to the overall performance, so it's fine to use reflection when it makes the program clearer. Testing is a particularly good fit for reflection since most tests use small data sets. But for functions on the critical path, reflection is best avoided.
