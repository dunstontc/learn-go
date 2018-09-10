# Chapter 7: Interfaces

<!-- TOC -->

- [7.1. Interfaces as Contracts](#71-interfaces-as-contracts)
- [7.2. Interface Types](#72-interface-types)
- [7.3. Interface Satisfaction](#73-interface-satisfaction)
- [7.4. Parsing Flags with flag.Value](#74-parsing-flags-with-flagvalue)
- [7.5. Interface Values](#75-interface-values)
- [7.6. Sorting with sort.Interface](#76-sorting-with-sortinterface)
- [7.7. The http.Handler Interface](#77-the-httphandler-interface)
- [7.8. The error Interface](#78-the-error-interface)
- [7.9. Example: Expression Evaluator](#79-example-expression-evaluator)
- [7.10. Type Assertions](#710-type-assertions)
- [7.11. Discriminating Errors with Type Assertions](#711-discriminating-errors-with-type-assertions)
- [7.12. Querying Behaviors with Interface Type Assertions](#712-querying-behaviors-with-interface-type-assertions)
- [7.13. Type Switches](#713-type-switches)
- [7.14. Example: Token-Based XML Decoding](#714-example-token-based-xml-decoding)
- [7.15. A Few Words of Advice](#715-a-few-words-of-advice)

<!-- /TOC -->

Interface types express generalizations or abstractions about the behaviors of other types. By generalizing, interfaces let us write functions that are more flexible and adaptable because they are not tied to the details of one particular implementation.

Many object-oriented languages have some notion of interfaces, but what makes Go's interfaces so distinctive is that they are *satisfied implicitly*. In other words, there's no need to declare all the interfaces that a given concrete type satisfies; simply possessing the necessary methods is enough. This design lets you create new interfaces that are satisfied by existing concrete types without changing the existing types, which is particularly useful for types defined in packages that you don't control.

In this chapter, we'll start by looking at the basic mechanics of interface types and their values. Along the way, we'll study several important interfaces from the standard library. Many Go programs make as much use of standard interfaces as they do of their own ones. Finally, we'll look at *type assertions* (§7.10) and *type switches* (§7.13) and see how they enable a different kind of generality.


## 7.1. Interfaces as Contracts

All the types we've looked at so far have been *concrete types*. A concrete type specifies the exact representation of its values and exposes the intrinsic operations of that representation, such as arithmetic for numbers, or indexing, append, and range for slices. A concrete type may also provide additional behaviors through its methods. When you have a value of a concrete type, you know exactly what it is and what you can do with it.

There is another kind of type in Go called an *interface type*. An interface is an *abstract type*. It doesn't expose the representation or internal structure of its values, or the set of basic operations they support; it reveals only some of their methods. When you have a value of an interface type, you know nothing about what it is; you know only what it can do, or more precisely, what behaviors are provided by its methods.

Throughout the book, we've been using two similar functions for string formatting: `fmt.Printf`, which writes the result to the standard output (a file), and `fmt.Sprintf`, which returns the result as a string. It would be unfortunate if the hard part, formatting the result, had to be duplicated because of these superficial differences in how the result is used. Thanks to interfaces, it does not. Both of these functions are, in effect, wrappers around a third function, `fmt.Fprintf`, that is agnostic about what happens to the result it computes:
```go
  package fmt

  func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)

  func Printf(format string, args ...interface{}) (int, error) {
      return Fprintf(os.Stdout, format, args...)
  }

  func Sprintf(format string, args ...interface{}) string {
      var buf bytes.Buffer
      Fprintf(&buf, format, args...)
      return buf.String()
  }
```

The `F` prefix of `Fprintf` stands for file and indicates that the formatted output should be written to the file provided as the first argument. In the `Printf` case, the argument, `os.Stdout,` is an `*os.File`. In the `Sprintf` case, however, the argument is not a file, though it superficially resembles one: `&buf` is a pointer to a memory buffer to which bytes can be written.

The first parameter of `Fprintf` is not a file either. It's an `io.Writer`, which is an interface type with the following declaration:
```go
  package io

  // Writer is the interface that wraps the basic Write method.
  type Writer interface {
      // Write writes len(p) bytes from p to the underlying data stream.
      // It returns the number of bytes written from p (0 <= n <= len(p))
      // and any error encountered that caused the write to stop early.
      // Write must return a non-nil error if it returns n < len(p).
      // Write must not modify the slice data, even temporarily.
      //
      // Implementations must not retain p.
      Write(p []byte) (n int, err error)
  }
```
The `io.Writer` interface defines the contract between Fprintf and its callers. On the one hand, the contract requires that the caller provide a value of a concrete type like `*os.File` or `*bytes.Buffer` that has a method called `Write` with the appropriate signature and behavior. On the other hand, the contract guarantees that `Fprintf` will do its job given any value that satisfies the `io.Writer` interface. `Fprintf` may not assume that it is writing to a file or to memory, only that it can call Write.

Because `fmt.Fprintf` assumes nothing about the representation of the value and relies only on the behaviors guaranteed by the io.Writer contract, we can safely pass a value of any concrete type that satisfies `io.Writer` as the first argument to `fmt.Fprintf`. This freedom to substitute one type for another that satisfies the same interface is called *substitutability*, and is a hallmark of object-oriented programming.

Let's test this out using a new type. The `Write` method of the `*ByteCounter` type below merely counts the bytes written to it before discarding them. (The conversion is required to make the types of `len(p)` and `*c` match in the `+=` assignment statement.)
```go
// gopl.io/ch7/bytecounter
// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}
```
Since `*ByteCounter` satisfies the `io.Writer` contract, we can pass it to `Fprintf`, which does its string formatting oblivious to this change; the `ByteCounter` correctly accumulates the length of the result.
```go
    var c ByteCounter
    c.Write([]byte("hello"))
    fmt.Println(c) // "5", = len("hello")
    c = 0 // reset the counter
    
    var name = "Dolly"
    fmt.Fprintf(&c, "hello, %s", name)
    fmt.Println(c) // "12", = len("hello, Dolly")
```
Besides io.Writer, there is another interface of great importance to the `fmt` package. `Fprintf` and `Fprintln` provide a way for types to control how their values are printed. In Section 2.5, we defined a `String` method for the `Celsius` type so that temperatures would print as "100°C", and in Section 6.5 we equipped `*IntSet` with a `String` method so that sets would be rendered using traditional set notation like "{1 2 3}". Declaring a `String` method makes a type satisfy one of the most widely used interfaces of all, `fmt.Stringer`:
```go
    package fmt

    // The String method is used to print values passed
    // as an operand to any format that accepts a string
    // or to an unformatted printer such as Print.
    type Stringer interface {
        String() string
    }
```
We'll explain how the fmt package discovers which values satisfy this interface in Section 7.10.

### Exercises
- **Exercise 7.1**: Using the ideas from `ByteCounter`, implement counters for words and for lines. You will find `bufio.ScanWords` useful.
- **Exercise 7.2**: Write a function `CountingWriter` with the signature below that, given an `io.Writer`, returns a new Writer that wraps the original, and a pointer to an int64 variable that at any moment contains the number of bytes written to the new `Writer`.
```go
  func CountingWriter(w io.Writer) (io.Writer, *int64)
```
- **Exercise 7.3**: Write a `String` method for the `*tree` type in `gopl.io/ch4/treesort` (§4.4) that reveals the sequence of values in the tree.


## 7.2. Interface Types 

An interface type specifies a set of methods that a concrete type must possess to be considered an instance of that interface.

The io.Writer type is one of the most widely used interfaces because it provides an abstraction of all the types to which bytes can be written, which includes files, memory buffers, network connections, HTTP clients, archivers, hashers, and so on. The `io` package defines many other useful interfaces. A `Reader` represents any type from which you can read bytes, and a `Closer` is any value that you can close, such as a file or a network connection. (By now you've probably noticed the naming convention for many of Go's single-method interfaces.)
```go
  package io

  type Reader interface {
      Read(p []byte) (n int, err error)
  }
      
  type Closer interface {
      Close() error
  }
```
Looking farther, we find declarations of new interface types as combinations of existing ones. Here are two examples:
```go
  type ReadWriter interface {
      Reader
      Writer
  }
      
  type ReadWriteCloser interface {
      Reader
      Writer
      Closer
  }
```
The syntax used above, which resembles struct embedding, lets us name another interface as a shorthand for writing out all of its methods. This is called *embedding* an interface. We could have written `io.ReadWriter` without embedding, albeit less succinctly, like this:
```go
  type ReadWriter interface {
      Read(p []byte) (n int, err error)
      Write(p []byte) (n int, err error)
  }
```
or even using a mixture of the two styles:
```go
type ReadWriter interface {
         Read(p []byte) (n int, err error)
         Writer
}
```


### Exercises
- **Exercise 7.4**: The `strings.NewReader` function returns a value that satisfies the `io.Reader` interface (and others) by reading from its argument, a string. Implement a simple version of `NewReader` yourself, and use it to make the HTML parser (§5.2) take input from a string.
- **Exercise 7.5**: The LimitReader function in the io package accepts an `io.Reader` `r` and `a` number of bytes `n`, and returns another `Reader` that reads from `r` but reports an end-of-file condition after `n` bytes. Implement it.
```go
  func LimitReader(r io.Reader, n int64) io.Reader
```


## 7.3. Interface Satisfaction 

A type *satisfies* an interface if it possesses all the methods the interface requires. For example, an `*os.File` satisfies `io.Reader`,`Writer`, `Closer`, and `ReadWriter`. A `*bytes.Buffer` stisfies `Reader`, `Writer`, and `ReadWriter`, but does not satisfy `Closer` because it does not have a `Close` method. As a shorthand, Go programmers often say that a concrete type *"is a"* particular interface type, meaning that it satisfies the interface. For example, a `*bytes.Buffer` is an `io.Writer`; an `*os.File` is an `io.ReadWriter`.

The assignability rule (§2.4.2) for interfaces is very simple: an expression may be assigned to an interface only if its type satisfies the interface. So:
```go
    var w io.Writer
    w = os.Stdout           // OK: *os.File has Write method
    w = new(bytes.Buffer)   // OK: *bytes.Buffer has Write method
    w = time.Second         // compile error: time.Duration lacks Write method
    
    var rwc io.ReadWriteCloser
    rwc = os.Stdout         // OK: *os.File has Read, Write, Close methods
    rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
```
This rule applies even when the right-hand side is itself an interface:
```go
    w = rwc                 // OK: io.ReadWriteCloser has Write method
    rwc = w                 // compile error: io.Writer lacks Close method
```
Because `ReadWriter` and `ReadWriteCloser` include all the methods of `Writer`, any type that satisfies `ReadWriter` or `ReadWriteCloser` necessarily satisfies `Writer`.

Before we go further, we should explain one subtlety in what it means for a type to have a method. Recall from Section 6.2 that for each named concrete type `T`, some of its methods have a receiver of type `T` itself whereas others require a `*T` pointer. Recall also that it is legal to call a `*T` method on an argument of type `T` so long as the argument is a variable; the compiler implicitly takes its address. But this is mere syntactic sugar: a value of type `T` does not possess all the methods that a `*T` pointer does, and as a result it might satisfy fewer interfaces.

An example will make this clear. The `String` method of the `IntSet` type from Section 6.5 requires a pointer receiver, so we cannot call that method on a non-addressable `IntSet` value:
```go
    type IntSet struct { /* ... */ }
    func (*IntSet) String() string
    
    var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
```
but we can call it on an `IntSet` variable:
```go
    var s IntSet
    var _ = s.String() // OK: s is a variable and &s has a String method
```
However, since only `*IntSet` has a `String` method, only `*IntSet` satisfies the `fmt.Stringer` interface:
```go
    var _ fmt.Stringer = &s // OK
    var _ fmt.Stringer = s  // compile error: IntSet lacks String method
```
Section 12.8 includes a program that prints the methods of an arbitrary value, and the `godoc -analysis=type` tool (§10.7.4) displays the methods of each type and the relationship between interfaces and concrete types.

Like an envelope that wraps and conceals the letter it holds, an interface wraps and conceals the concrete type and value that it holds. Only the methods revealed by the interface type may be called, even if the concrete type has others:
```go
    os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method
    os.Stdout.Close()                // OK: *os.File has Close method

    var w io.Writer
    w = os.Stdout
    w.Write([]byte("hello")) // OK: io.Writer has Write method
    w.Close()                // compile error: io.Writer lacks Close method
```

An interface with more methods, such as `io.ReadWriter`, tells us more about the values it contains, and places greater demands on the types that implement it, than does an interface with fewer methods such as `io.Reader`. So what does the type `interface{}`, which has no methods at all, tell us about the concrete types that satisfy it?

That's right: nothing. This may seem useless, but in fact the type `interface{}`, which is called the *empty interface* type, is indispensable. Because the empty interface type places no demands on the types that satisfy it, we can assign *any* value to the empty interface.
```go
    var any interface{}
    any = true
    any = 12.34
    any = "hello"
    any = map[string]int{"one": 1}
    any = new(bytes.Buffer)
```
Although it wasn't obvious, we've been using the empty interface type since the very first example in this book, because it is what allows functions like `fmt.Println`, or `errorf` in Section 5.7, to accept arguments of any type.

Of course, having created an `interface{}` value containing a boolean, float, string, map, pointer, or any other type, we can do nothing directly to the value it holds since the interface has no methods. We need a way to get the value back out again. We'll see how to do that using a *type assertion* in Section 7.10.

Since interface satisfaction depends only on the methods of the two types involved, there is no need to declare the relationship between a concrete type and the interfaces it satisfies. That said, it is occasionally useful to document and assert the relationship when it is intended but not otherwise enforced by the program. The declaration below asserts at compile time that a value of type `*bytes.Buffer` satisfies `io.Writer`:
```go
    // *bytes.Buffer must satisfy io.Writer
    var w io.Writer = new(bytes.Buffer)
```
We needn't allocate a new variable since any value of type `*bytes.Buffer` will do, even `nil`, which we write as `(*bytes.Buffer)(nil)` using an explicit conversion. And since we never intend to refer to `w`, we can replace it with the blank identifier. Together, these changes give us this more frugal variant:
```go
    // *bytes.Buffer must satisfy io.Writer
    var _ io.Writer = (*bytes.Buffer)(nil)
```
Non-empty interface types such as `io.Writer` are most often satisfied by a pointer type, particularly when one or more of the interface methods implies some kind of mutation to the receiver, as the `Write` method does. A pointer to a struct is an especially common method-bearing type.

But pointer types are by no means the only types that satisfy interfaces, and even interfaces with mutator methods may be satisfied by one of Go's other reference types. We've seen examples of slice types with methods (`geometry.Path`, §6.1) and map types with methods (`url.Values`, §6.2.1), and later we'll see a function type with methods (`http.HandlerFunc`, §7.7). Even basic types may satisfy interfaces; as we saw in Section 7.4, `time.Duration` satisfies `fmt.Stringer`.

A concrete type may satisfy many unrelated interfaces. Consider a program that organizes or sells digitized cultural artifacts like music, films, and books. It might define the following set of concrete types:
```
    Album
    Book
    Movie
    Magazine
    Podcast
    TVEpisode
    Track
```
We can express each abstraction of interest as an interface. Some properties are common to all artifacts, such as a title, a creation date, and a list of creators (authors or artists).
```go
    type Artifact interface {
        Title() string
        Creators() []string
        Created() time.Time
    }
```
Other properties are restricted to certain types of artifacts. Properties of the printed word are relevant only to books and magazines, whereas only movies and TV episodes have a screen resolution.
```go
    type Text interface {
        Pages() int
        Words() int
        PageSize() int
    }

    type Audio interface {
        Stream() (io.ReadCloser, error)
        RunningTime() time.Duration
        Format() string // e.g., "MP3", "WAV"
    }

    type Video interface {
        Stream() (io.ReadCloser, error)
        RunningTime() time.Duration
        Format() string // e.g., "MP4", "WMV"
        Resolution() (x, y int)
    }
```
These interfaces are but one useful way to group related concrete types together and express the facets they share in common. We may discover other groupings later. For example, if we find we need to handle `Audio` and `Video` items in the same way, we can define a `Streamer` interface to represent their common aspects without changing any existing type declarations.
```go
    type Streamer interface {
        Stream() (io.ReadCloser, error)
        RunningTime() time.Duration
        Format() string
    }
```
Each grouping of concrete types based on their shared behaviors can be expressed as an interface type. Unlike class-based languages, in which the set of interfaces satisfied by a class is explicit, in Go we can define new abstractions or groupings of interest when we need them, without modifying the declaration of the concrete type. This is particularly useful when the concrete type comes from a package written by a different author. Of course, there do need to be underlying commonalities in the concrete types.


## 7.4. Parsing Flags with flag.Value 

In this section, we'll see how another standard interface, `flag.Value`, helps us define new notations for command-line flags. Consider the program below, which sleeps for a specified period of time.
```go
// gopl.io/ch7/sleep
// The sleep program sleeps for a specified period of time.
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
```
Before it goes to sleep it prints the time period. The `fmt` package calls the `time.Duration`'s `String` method to print the period not as a number of nanoseconds, but in a user-friendly notation:
```
    $ go build gopl.io/ch7/sleep
    $ ./sleep
    Sleeping for 1s...
```
By default, the sleep period is one second, but it can be controlled through the `-period` command-line flag. The `flag.Duration` function creates a flag variable of type `time.Duration` and allows the user to specify the duration in a variety of user-friendly formats, including the same notation printed by the `String` method. This symmetry of design leads to a nice user interface.
```
    $ ./sleep -period 50ms
    Sleeping for 50ms...
    $ ./sleep -period 2m30s
    Sleeping for 2m30s...
    $ ./sleep -period 1.5h
    Sleeping for 1h30m0s...
    $ ./sleep -period "1 day"
    invalid value "1 day" for flag -period: time: invalid duration 1 day
```
Because duration-valued flags are so useful, this feature is built into the flag package, but it's easy to define new flag notations for our own data types. We need only define a type that satisfies the `flag.Value` interface, whose declaration is below:
```go
    package flag

    // Value is the interface to the value stored in a flag.
    type Value interface {
        String() string
        Set(string) error
    }
```
The `String` method formats the flag's value for use in command-line help messages; thus every `flag.Value` is also a `fmt.Stringer`. The `Set` method parses its string argument and updates the flag value. In effect, the `Set` method is the inverse of the `String` method, and it is good practice for them to use the same notation.

Let's define a `celsiusFlag` type that allows a temperature to be specified in Celsius, or in Fahrenheit with an appropriate conversion. Notice that `celsiusFlag` embeds a `Celsius` (§2.5), thereby getting a `String` method for free. To satisfy `flag.Value`, we need only declare the `Set` method:
```go
// gopl.io/ch7/tempconv
// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
```
The call to `fmt.Sscanf` parses a floating-point number (`value`) and a string (`unit`) from the input `s`. Although one must usually check `Sscanf`'s error result, in this case we don't need to because if there was a problem, no switch case will match.

The `CelsiusFlag` function below wraps it all up. To the caller, it returns a pointer to the `Celsius` field embedded within the `celsiusFlag` variable `f`. The `Celsius` field is the variable that will be updated by the `Set` method during flags processing. The call to `Var` adds the flag to the application's set of command-line flags, the global variable `flag.CommandLine`. Programs with unusually complex command-line interfaces may have several variables of this type. The call to `Var` assigns a `*celsiusFlag` argument to a `flag.Value` parameter, causing the compiler to check that `*celsiusFlag` has the necessary methods.
```go
// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
```
Now we can start using the new flag in our programs:
```go
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
```
Here's a typical session:
```
    $ go build gopl.io/ch7/tempflag
    $ ./tempflag
    20°C
    $ ./tempflag -temp -18C
    -18°C
    $ ./tempflag -temp 212°F
    100°C
    $ ./tempflag -temp 273.15K
    invalid value "273.15K" for flag -temp: invalid temperature "273.15K"
    Usage of ./tempflag:
      -temp value
            the temperature (default 20°C)
    $ ./tempflag -help
    Usage of ./tempflag:
      -temp value
            the temperature (default 20°C)
```

### Exercises
- **Exercise 7.6**: Add support for Kelvin temperatures to `tempflag`.
- **Exercise 7.7**: Explain why the help message contains `°C` when the default value of `20.0` does not.


## 7.5. Interface Values

Conceptually, a value of an interface type, or *interface value*, has two components, a concrete type and a value of that type. These are called the interface's *dynamic type* and *dynamic value*.

For a statically typed language like Go, types are a compile-time concept, so a type is not a value. In our conceptual model, a set of values called *type descriptors* provide information about each type, such as its name and methods. In an interface value, the type component is represented by the appropriate type descriptor.


In the four statements below, the variable `w` takes on three different values. (The initial and final values are the same.)
```go
    var w io.Writer
    w = os.Stdout
    w = new(bytes.Buffer)
    w = nil
```
Let's take a closer look at the value and dynamic behavior of `w` after each statement. The first statement declares `w`:
```go
    var w io.Writer
```
In Go, variables are always initialized to a well-defined value, and interfaces are no exception. The zero value for an interface has both its type and value components set to nil (Figure 7.1).

![Figure 7.1](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig7.1.png)

An interface value is described as nil or non-nil based on its dynamic type, so this is a nil interface value. You can test whether an interface value is nil using `w == nil` or `w != nil`. Calling any method of a nil interface value causes a panic:
```go
    w.Write([]byte("hello")) // panic: nil pointer dereference
```
The second statement assigns a value of type `*os.File` to `w`:
```go
    w = os.Stdout
```
This assignment involves an implicit conversion from a concrete type to an interface type, and is equivalent to the explicit conversion `io.Writer(os.Stdout)`. A conversion of this kind, whether explicit or implicit, captures the type and the value of its operand. The interface value's dynamic type is set to the type descriptor for the pointer type `*os.File`, and its dynamic value holds a copy of `os.Stdout`, which is a pointer to the `os.File` variable representing the standard output of the process (Figure 7.2).

![Figure 7.2](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig7.2.png)

Calling the Write method on an interface value containing an `*os.File` pointer causes the `(*os.File).Write` method to be called. The call prints `"hello"`.
```go
    w.Write([]byte("hello")) // "hello"
```
In general, we cannot know at compile time what the dynamic type of an interface value will be, so a call through an interface must use *dynamic dispatch*. Instead of a direct call, the compiler must generate code to obtain the address of the method named Write from the type descriptor, then make an indirect call to that address. The receiver argument for the call is a copy of the interface's dynamic value, `os.Stdout`. The effect is as if we had made this call directly:
```go
    os.Stdout.Write([]byte("hello")) // "hello"
```
The third statement assigns a value of type `*bytes.Buffer` to the interface value:
```go
    w = new(bytes.Buffer)
```
The dynamic type is now `*bytes.Buffer` and the dynamic value is a pointer to the newly allocated buffer (Figure 7.3).

![Figure 7.3](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig7.3.png)

A call to the Write method uses the same mechanism as before:
```go
    w.Write([]byte("hello")) // writes "hello" to the bytes.Buffer
```
This time, the type descriptor is `*bytes.Buffer`, so the `(*bytes.Buffer).Write` method is called, with the address of the buffer as the value of the receiver parameter. The call appends `"hello"` to the buffer.

Finally, the fourth statement assigns nil to the interface value:
```go
    w = nil
```
This resets both its components to `nil`, restoring `w` to the same state as when it was declared, which was shown in Figure 7.1.

An interface value can hold arbitrarily large dynamic values. For example, the `time.Time` type, which represents an instant in time, is a struct type with several unexported fields. If we create an interface value from it,
```go
    var x interface{} = time.Now()
```
the result might look like Figure 7.4. Conceptually, the dynamic value always fits inside the interface value, no matter how large its type. (This is only a conceptual model; a realistic implementation is quite different.)

![Figure 7.4](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig7.4.png)

Interface values may be compared using `==` and `!=`. Two interface values are equal if both are nil, or if their dynamic types are identical and their dynamic values are equal according to the usual behavior of `==` for that type. Because interface values are comparable, they may be used as the keys of a map or as the operand of a `switch` statement.

However, if two interface values are compared and have the same dynamic type, but that type is not comparable (a slice, for instance), then the comparison fails with a panic:
```go
    var x interface{} = []int{1, 2, 3}
    fmt.Println(x == x) // panic: comparing uncomparable type []int
```
In this respect, interface types are unusual. Other types are either safely comparable (like basic types and pointers) or not comparable at all (like slices, maps, and functions), but when comparing interface values or aggregate types that contain interface values, we must be aware of the potential for a panic. A similar risk exists when using interfaces as map keys or switch operands. Only compare interface values if you are certain that they contain dynamic values of comparable types.

When handling errors, or during debugging, it is often helpful to report the dynamic type of an interface value. For that, we use the fmt package's `%T` verb:
```go
    var w io.Writer
    fmt.Printf("%T\n", w) // "<nil>"

    w = os.Stdout
    fmt.Printf("%T\n", w) // "*os.File"

    w = new(bytes.Buffer)
    fmt.Printf("%T\n", w) // "*bytes.Buffer"
```
Internally, `fmt` uses reflection to obtain the name of the interface's dynamic type. We'll look at reflection in Chapter 12.


### 7.5.1. Caveat: An Interface Containing a Nil Pointer Is Non-Nil

A nil interface value, which contains no value at all, is not the same as an interface value containing a pointer that happens to be nil. This subtle distinction creates a trap into which every Go programmer has stumbled.

Consider the program below. With `debug` set to `true`, the main function collects the output of the function `f` in a `bytes.Buffer`.
```go
    const debug = true

    func main() {
        var buf *bytes.Buffer
        if debug {
            buf = new(bytes.Buffer) // enable collection of output
        }
        f(buf) // NOTE: subtly incorrect!
        if debug {
            // ...use buf...
        } 
    }

    // If out is non-nil, output will be written to it.
    func f(out io.Writer) {
        // ...do something...
        if out != nil {
            out.Write([]byte("done!\n"))
        } 
    }
```
We might expect that changing `debug` to `false` would disable the collection of the output, but in fact it causes the program to panic during the `out.Write` call:
```go
    if out != nil {
        out.Write([]byte("done!\n")) // panic: nil pointer dereference
    }
```
When `main` calls `f`, it assigns a nil pointer of type `*bytes.Buffer` to the out parameter, so the dynamic value of `out` is `nil`. However, its dynamic type is `*bytes.Buffer`, meaning that `out` is a non-nil interface containing a nil pointer value (Figure 7.5), so the defensive check `out != nil` is still true.

![Figure 7.5](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig7.5.png)

As before, the dynamic dispatch mechanism determines that `(*bytes.Buffer).Write` must be called but this time with a receiver value that is nil. For some types, such as `*os.File`, `nil` is a valid receiver (§6.2.1), but `*bytes.Buffer` is not among them. The method is called, but it panics as it tries to access the buffer.

The problem is that although a nil `*bytes.Buffer` pointer has the methods needed to satisfy the interface, it doesn't satisfy the behavioral requirements of the interface. In particular, the call violates the implicit precondition of `(*bytes.Buffer).Write` that its receiver is not nil, so assigning the nil pointer to the interface was a mistake. The solution is to change the type of `buf` in `main` to `io.Writer`, thereby avoiding the assignment of the dysfunctional value to the interface in the first place:
```go
    var buf io.Writer
    if debug {
        buf = new(bytes.Buffer) // enable collection of output
    }
    f(buf) // OK
```
Now that we've covered the mechanics of interface values, let's take a look at some more important interfaces from Go's standard library. In the next three sections, we'll see how interfaces are used for sorting, web serving, and error handling.


## 7.6. Sorting with `sort.Interface` 

Like string formatting, sorting is a frequently used operation in many programs. Although a minimal Quicksort can be written in about 15 lines, a robust implementation is much longer, and it is not the kind of code we should wish to write anew or copy each time we need it.

Fortunately, the `sort` package provides in-place sorting of any sequence according to any ordering function. Its design is rather unusual. In many languages, the sorting algorithm is associated with the sequence data type, while the ordering function is associated with the type of the elements. By contrast, Go's `sort.Sort` function assumes nothing about the representation of either the sequence or its elements. Instead, it uses an interface, `sort.Interface`, to specify the contract between the generic sort algorithm and each sequence type that may be sorted. An implementation of this interface determines both the concrete representation of the sequence, which is often a slice, and the desired ordering of its elements.

An in-place sort algorithm needs three things:
1. The length of the sequence
2. A means of comparing two elements
2. A way to swap two elements

so they are the three methods of `sort.Interface`:
```go
    package sort
    type Interface interface {
        Len() int
        Less(i, j int) bool // i, j are indices of sequence elements
        Swap(i, j int)
    }
```
To sort any sequence, we need to define a type that implements these three methods, then apply `sort.Sort` to an instance of that type. As perhaps the simplest example, consider sorting a slice of strings. The new type `StringSlice` and its `Len`, `Less`,and `Swap` methods are shown below.
```go
    type StringSlice []string
    func (p StringSlice) Len() int           { return len(p) }
    func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
    func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
```
Now we can sort a slice of strings, `names`, by converting the slice to a `StringSlice` like this:
```go
    sort.Sort(StringSlice(names))
```
The conversion yields a slice value with the same length, capacity, and underlying array as `names` but with a type that has the three methods required for sorting.

Sorting a slice of strings is so common that the `sort` package provides the `StringSlice` type, as well as a function called `Strings` so that the call above can be simplified to `sort.Strings(names)`.

The technique here is easily adapted to other sort orders, for instance, to ignore capitalization or special characters. (The Go program that sorts index terms and page numbers for this book does this, with extra logic for Roman numerals.) For more complicated sorting, we use the same idea, but with more complicated data structures or more complicated implementations of the `sort.Interface` methods.

Our running example for sorting will be a music playlist, displayed as a table. Each track is a single row, and each column is an attribute of that track, like artist, title, and running time. Imagine that a graphical user interface presents the table, and that clicking the head of a column causes the playlist to be sorted by that attribute; clicking the same column head again reverses the order. Let's look at what might happen in response to each click.

The variable `tracks` below contains a playlist. (One of the authors apologizes for the other author's musical tastes.) Each element is indirect, a pointer to a `Track`. Although the code below would work if we stored the `Track`s directly, the sort function will swap many pairs of elements, so it will run faster if each element is a pointer, which is a single machine word, instead of an entire `Track`, which might be eight words or more.
```go
// gopl.io/ch7/sorting
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
```
The `printTracks` function prints the playlist as a table. A graphical display would be nicer, but this little routine uses the `text/tabwriter` package to produce a table whose columns are neatly aligned and padded as shown below. Observe that `*tabwriter.Writer` satisfies `io.Writer`. It collects each piece of data written to it; its `Flush` method formats the entire table and writes it to `os.Stdout`.
```go
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}
```
To sort the playlist by the `Artist` field, we define a new slice type with the necessary `Len`, `Less`, and `Swap` methods, analogous to what we did for `StringSlice`.
```go
    type byArtist []*Track

    func (x byArtist) Len() int           { return len(x) }
    func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
    func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
```
To call the generic sort routine, we must first convert tracks to the new type, `byArtist`, that defines the order:
```go
    sort.Sort(byArtist(tracks))
```
After sorting the slice by artist, the output from `printTracks` is
```
    Title       Artist          Album              Year  Length
    -----       ------          -----              ----  ------
    Go Ahead    Alicia Keys     As I Am            2007  4m36s
    Go          Delilah         From the Roots Up  2012  3m38s
    Ready 2 Go  Martin Solveig  Smash              2011  4m24s
    Go          Moby            Moby               1992  3m37s
```
If the user requests "sort by artist" a second time, we'll sort the tracks in reverse. We needn't define a new type `byReverseArtist` with an inverted `Less` method, however, since the `sort` package provides a `Reverse` function that transforms any sort order to its inverse.
```go
    sort.Sort(sort.Reverse(byArtist(tracks)))
```
After reverse-sorting the slice by artist, the output from `printTracks` is
```
    Title       Artist          Album              Year  Length
    -----       ------          -----              ----  ------
    Go          Moby            Moby               1992  3m37s
    Ready 2 Go  Martin Solveig  Smash              2011  4m24s
    Go          Delilah         From the Roots Up  2012  3m38s
    Go Ahead    Alicia Keys     As I Am            2007  4m36s
```
The `sort.Reverse` function deserves a closer look since it uses composition (§6.3), which is an important idea. The `sort` package defines an unexported type `reverse`, which is a struct that embeds a `sort.Interface`. The `Less` method for `reverse` calls the `Less` method of the embedded `sort.Interface` value, but with the indices flipped, reversing the order of the sort results.
```go
    package sort

    type reverse struct{ Interface } // that is, sort.Interface

    func (r reverse) Less(i, j int) bool { return r.Interface.Less(j, i) }

    func Reverse(data Interface) Interface { return reverse{data} }
```
`Len` and `Swap`, the other two methods of reverse, are implicitly provided by the original `sort.Interface` value because it is an embedded field. The exported function `Reverse` returns an instance of the `reverse` type that contains the original `sort.Interface` value.

To sort by a different column, we must define a new type, such as `byYear`:
```go
    type byYear []*Track

    func (x byYear) Len() int           { return len(x) }
    func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
    func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
```
After sorting tracks by year using `sort.Sort(byYear(tracks))`, `printTracks` shows a chronological listing:
```
    Title       Artist          Album              Year  Length
    -----       ------          -----              ----  ------
    Go          Moby            Moby               1992  3m37s
    Go Ahead    Alicia Keys     As I Am            2007  4m36s
    Ready 2 Go  Martin Solveig  Smash              2011  4m24s
    Go          Delilah         From the Roots Up  2012  3m38s
```
For every slice element type and every ordering function we need, we declare a new implementation of `sort.Interface`. As you can see, the `Len` and `Swap` methods have identical definitions for all slice types. In the next example, the concrete type `customSort` combines a slice with a function, letting us define a new sort order by writing only the comparison function. Incidentally, the concrete types that implement `sort.Interface` are not always slices; `customSort` is a struct type.
```go
    type customSort struct {
        t    []*Track
        less func(x, y *Track) bool
    }

    func (x customSort) Len() int           { return len(x.t) }
    func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
    func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
```
Let's define a multi-tier ordering function whose primary sort key is the `Title`, whose secondary key is the `Year`, and whose tertiary key is the running time, `Length`. Here's the call to `Sort` using an anonymous ordering function:
```go
    sort.Sort(customSort{tracks, func(x, y *Track) bool {
        if x.Title != y.Title {
            return x.Title < y.Title
        }
        if x.Year != y.Year {
            return x.Year < y.Year
        }
        if x.Length != y.Length {
            return x.Length < y.Length
        }
        return false
    }})
```
And here's the result. Notice that the tie between the two tracks titled "Go" is broken in favor of the older one.
```
    Title       Artist          Album              Year  Length
    -----       ------          -----              ----  ------
    Go          Moby            Moby               1992  3m37s
    Go          Delilah         From the Roots Up  2012  3m38s
    Go Ahead    Alicia Keys     As I Am            2007  4m36s
    Ready 2 Go  Martin Solveig  Smash              2011  4m24s
```
Although sorting a sequence of length *n* requires O(*n* log *n*) comparison operations, testing whether a sequence is already sorted requires at most *n−1* comparisons. The `IsSorted` function from the sort package checks this for us. Like `sort.Sort`, it abstracts both the sequence and its ordering function using `sort.Interface`, but it never calls the `Swap` method: This code demonstrates the `IntsAreSorted` and `Ints` functions and the `IntSlice` type:
```go
    values := []int{3, 1, 4, 1}
    fmt.Println(sort.IntsAreSorted(values)) // "false"
    sort.Ints(values)
    fmt.Println(values)                     // "[1 1 3 4]"
    fmt.Println(sort.IntsAreSorted(values)) // "true"
    sort.Sort(sort.Reverse(sort.IntSlice(values)))
    fmt.Println(values)                     // "[4 3 1 1]"
    fmt.Println(sort.IntsAreSorted(values)) // "false"
```
For convenience, the `sort` package provides versions of its functions and types specialized for `[]int`, `[]string`, and `[]float64` using their natural orderings. For other types, such as `[]int64` or `[]uint`, we're on our own, though the path is short.


### Exercises
- **Exercise 7.8**: Many GUIs provide a table widget with a stateful multi-tier sort: the primary sort key is the most recently clicked column head, the secondary sort key is the second-most recently clicked column head, and so on. Define an implementation of `sort.Interface` for use by such a table. Compare that approach with repeated sorting using `sort.Stable`.
- **Exercise 7.9**: Use the `html/template` package (§4.6) to replace `printTracks` with a function that displays the `tracks` as an HTML table. Use the solution to the previous exercise to arrange that each click on a column head makes an HTTP request to sort the table.
- **Exercise 7.10**: The `sort.Interface` type can be adapted to other uses. Write a function `IsPalindrome(s sort.Interface) bool` that reports whether the sequence `s` is a palindrome, in other words, reversing the sequence would not change it. Assume that the elements at indices `i` and `j` are equal if `!s.Less(i, j) && !s.Less(j, i)`.


## 7.7. The `http.Handler` Interface 

In Chapter 1, we saw a glimpse of how to use the `net/http` package to implement web clients (§1.5) and servers (§1.7). In this section, we'll look more closely at the server API, whose foundation is the http.Handler interface:
```go
    // net/http
    package http
       
    type Handler interface {
        ServeHTTP(w ResponseWriter, r *Request)
    }

    func ListenAndServe(address string, h Handler) error
```
The `ListenAndServe` function requires a server address, such as `"localhost:8000"`, and an instance of the `Handler` interface to which all requests should be dispatched. It runs forever, or until the server fails (or fails to start) with an error, always non-nil, which it returns.

Imagine an e-commerce site with a database mapping the items for sale to their prices in dol- lars. The program below shows the simplest imaginable implementation. It models the inventory as a map type, `database`, to which we've attached a `ServeHTTP` method so that it satisfies the `http.Handler` interface. The handler ranges over the map and prints the items.
```go
// gopl.io/ch7/http1
func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
```
If we start the server,
```bash
    $ go build gopl.io/ch7/http1
    $ ./http1 &
```
then connect to it with the `fetch` program from Section 1.5 (or a web browser if you prefer), we get the following output:
```bash
    $ go build gopl.io/ch1/fetch
    $ ./fetch http://localhost:8000
    shoes: $50.00
    socks: $5.00
```
So far, the server can only list its entire inventory and will do this for every request, regardless of URL. A more realistic server defines multiple different URLs, each triggering a different behavior. Let's call the existing one `/list` and add another one called `/price` that reports the price of a single item, specified as a request parameter like `/price?item=socks`.
```go
// gopl.io/ch7/http2
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}
```
Now the handler decides what logic to execute based on the path component of the URL, `req.URL.Path`. If the handler doesn't recognize the path, it reports an HTTP error to the client by calling `w.WriteHeader(http.StatusNotFound)`; this must be done before writing any text to `w`. (Incidentally, `http.ResponseWriter` is another interface. It augments `io.Writer` with methods for sending HTTP response headers.) Equivalently, we could use the `http.Error` utility function:
```go
    msg := fmt.Sprintf("no such page: %s\n", req.URL)
    http.Error(w, msg, http.StatusNotFound) // 404
```
The case for `/price` calls the URL's Query method to parse the HTTP request parameters as a map, or more precisely, a  multimap of type `url.Values` (§6.2.1) from the `net/url` package. It then finds the first item parameter and looks up its price. If the item wasn't found, it reports an error.

Here's an example session with the new server:
```bash
    $ go build gopl.io/ch7/http2
    $ go build gopl.io/ch1/fetch
    $ ./http2 &
    $ ./fetch http://localhost:8000/list
    shoes: $50.00
    socks: $5.00
    $ ./fetch http://localhost:8000/price?item=socks
    $5.00
    $ ./fetch http://localhost:8000/price?item=shoes
    $50.00
    $ ./fetch http://localhost:8000/price?item=hat
    no such item: "hat"
    $ ./fetch http://localhost:8000/help
    no such page: /help
```
Obviously we could keep adding cases to `ServeHTTP`, but in a realistic application, it's convenient to define the logic for each case in a separate function or method. Furthermore, related URLs may need similar logic; several image files may have URLs of the form `/images/*.png`, for instance. For these reasons, `net/http` provides `ServeMux`, a *request multiplexer*, to simplify the association between URLs and handlers. A `ServeMux` aggregates a collection of `http.Handler`s into a single `http.Handler`. Again, we see that different types satisfying the same interface are *substitutable*: the web server can dispatch requests to any `http.Handler`, regardless of which concrete type is behind it.

For a more complex application, several `ServeMux`es may be composed to handle more intricate dispatching requirements. Go doesn't have a canonical web framework analogous to Ruby's Rails or Python's Django. This is not to say that such frameworks don't exist, but the building blocks in Go's standard library are flexible enough that frameworks are often unnecessary. Furthermore, although frameworks are convenient in the early phases of a project, their additional complexity can make longer-term maintenance harder.

In the program below, we create a `ServeMux` and use it to associate the URLs with the corresponding handlers for the `/list` and `/price` operations, which have been split into separate methods. We then use the `ServeMux` as the main handler in the call to `ListenAndServe`.
```go
// gopl.io/ch7/http3
func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
```
Let's focus on the two calls to `mux.Handle` that register the handlers. In the first one, `db.list` is a method value (§6.4), that is, a value of type
```go
    func(w http.ResponseWriter, req *http.Request)
```
that, when called, invokes the `database.list` method with the receiver value `db`. So `db.list` is a function that implements handler-like behavior, but since it has no methods, it doesn't satisfy the `http.Handler` interface and can't be passed directly to mux.Handle.

The expression `http.HandlerFunc(db.list)` is a conversion, not a function call, since `http.HandlerFunc` is a type. It has the following definition:
```go
  // net/http
  package http

  type HandlerFunc func(w ResponseWriter, r *Request)

  func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
      f(w, r)
  }
```
`HandlerFunc` demonstrates some unusual features of Go's interface mechanism. It is a function type that has methods and satisfies an interface, http.Handler. The behavior of its `ServeHTTP` method is to call the underlying function. `HandlerFunc` is thus an adapter that lets a function value satisfy an interface, where the function and the interface's sole method have the same signature. In effect, this trick lets a single type such as database satisfy the `http.Handler` interface several different ways: once through its `list` method, once through its `price` method, and so on.

Because registering a handler this way is so common, `ServeMux` has a convenience method called `HandleFunc` that does it for us, so we can simplify the handler registration code to this:
```go
// gopl.io/ch7/http3a
mux.HandleFunc("/list", db.list)
mux.HandleFunc("/price", db.price)
```
It's easy to see from the code above how one would construct a program in which there are two different web servers, listening on different ports, defining different URLs, and dispatching to different handlers. We would just construct another `ServeMux` and make another call to `ListenAndServe`, perhaps concurrently. But in most programs, one web server is plenty. Also, it's typical to define HTTP handlers across many files of an application, and it would be a nuisance if they all had to be explicitly registered with the application's `ServeMux` instance.

So, for convenience, `net/http` provides a global `ServeMux` instance called `DefaultServeMux` and package-level functions called `http.Handle` and `http.HandleFunc`. To use `DefaultServeMux` as the server's main handler, we needn't pass it to `ListenAndServe`; `nil` will do.

The server's `main` function can then be simplified to
```go
// gopl.io/ch7/http4
func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

```
Finally, an important reminder: as we mentioned in Section 1.7, the web server invokes each handler in a new goroutine, so handlers must take precautions such as *locking* when accessing variables that other goroutines, including other requests to the same handler, may be accessing. We'll talk about concurrency in the next two chapters.

### Exercises
- **Exercise 7.11**: Add additional handlers so that clients can create, read, update, and delete database entries. For example, a request of the form `/update?item=socks&price=6` will update the price of an item in the inventory and report an error if the item does not exist or if the price is invalid. (Warning: this change introduces concurrent variable updates.)
- **Exercise 7.12**: Change the handler for `/list` to print its output as an HTML table, not text. You may find the `html/template` package (§4.6) useful.


## 7.8. The `error` Interface 

Since the beginning of this book, we've been using and creating values of the mysterious predeclared `error` type without explaining what it really is. In fact, it's just an interface type with a single method that returns an error message:
```go
    type error interface {
        Error() string
    }
```
The simplest way to create an error is by calling `errors.New`, which returns a new `error` for a given error message. The entire `errors` package is only four lines long:
```go
    package errors
    
    func New(text string) error { return &errorString{text} }
    
    type errorString struct { text string }
    
    func (e *errorString) Error() string { return e.text }
```

The underlying type of `errorString` is a struct, not a string, to protect its representation from inadvertent (or premeditated) updates. And the reason that the pointer type `*errorString`, not `errorString` alone, satisfies the `error` interface is so that every call to `New` allocates a distinct `error` instance that is equal to no other. We would not want a distinguished `error` such as `io.EOF` to compare equal to one that merely happened to have the same message.
```go
    fmt.Println(errors.New("EOF") == errors.New("EOF")) // "false"
```
Calls to `errors.New` are relatively infrequent because there's a convenient wrapper function, `fmt.Errorf`, that does string formatting too. We used it several times in Chapter 5.
```go
    package fmt

    import "errors"

    func Errorf(format string, args ...interface{}) error {
        return errors.New(Sprintf(format, args...))
    }
```
Although `*errorString` may be the simplest type of `error`, it is far from the only one. For example, the `syscall` package provides Go's low-level system call API. On many platforms, it defines a numeric type `Errno` that satisfies `error`, and on Unix platforms, `Errno`'s `Error` method does a lookup in a table of strings, as shown below:
```go
    package syscall

    type Errno uintptr // operating system error code

    var errors = [...]string{
        1:   "operation not permitted",   // EPERM
        2:   "no such file or directory", // ENOENT
        3:   "no such process",           // ESRCH
        // ...
    }

    func (e Errno) Error() string {
        if 0 <= int(e) && int(e) < len(errors) {
            return errors[e]
        }
        return fmt.Sprintf("errno %d", e)
    }
```
The following statement creates an interface value holding the `Errno` value 2, signifying the POSIX `ENOENT` condition:
```go
    var err error = syscall.Errno(2)
    fmt.Println(err.Error()) // "no such file or directory"
    fmt.Println(err)         // "no such file or directory"
```
The value of err is shown graphically in Figure 7.6.

![Figure 7.6](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig7.6.png)

`Errno` is an efficient representation of system call errors drawn from a finite set, and it satisfies the standard `error` interface. We'll see other types that satisfy this interface in Section 7.11.


## 7.9. Example: Expression Evaluator 

In this section, we'll build an evaluator for simple arithmetic expressions. We'll use an interface, `Expr`, to represent any expression in this language. For now, this interface needs no methods, but we'll add some later.
```go
    // An Expr is an arithmetic expression.
    type Expr interface{}
```
Our expression language consists of floating-point literals; the binary operators `+`, `-`, `*`, and `/`; the unary operators `-x` and `+x`; function calls `pow(x,y)`, `sin(x)`, and `sqrt(x)`; variables such as `x` and `pi`; and of course parentheses and standard operator precedence. All values are of type `float64`. Here are some example expressions:
```go
    sqrt(A / pi)
    pow(x, 3) + pow(y, 3)
    (F - 32) * 5 / 9
```
The five concrete types below represent particular kinds of expression. A `Var` represents a reference to a variable. (We'll soon see why it is exported.) A `literal` represents a floating-point constant. The `unary` and `binary` types represent operator expressions with one or two operands, which can be any kind of `Expr`. A call represents a function `call`; we'll restrict its `fn` field to `pow`, `sin`, or `sqrt`.
```go
// gopl.io/ch7/eval
// A Var identifies a variable, e.g., x.
type Var string

// A literal is a numeric constant, e.g., 3.141.
type literal float64

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}
```
To evaluate an expression containing variables, we'll need an *environment* that maps variable names to values:
```go
type Env map[Var]float64
```
We'll also need each kind of expression to define an `Eval` method that returns the expression's value in a given environment. Since every expression must provide this method, we add it to the `Expr` interface. The package exports only the types `Expr`, `Env`, and `Var`; clients can use the evaluator without access to the other expression types.
```go
type Expr interface {
    // Eval returns the value of this Expr in the environment env.
    Eval(env Env) float64
}
```
The concrete `Eval` methods are shown below. The method for `Var` performs an environment lookup, which returns zero if the variable is not defined, and the method for `literal` simply returns the literal value.
```go
func (v Var) Eval(env Env) float64 {
  return env[v]
}

func (l literal) Eval(_ Env) float64 {
  return float64(l)
}
```
The `Eval` methods for `unary` and `binary` recursively evaluate their operands, then apply the operation `op` to them. We don't consider divisions by zero or infinity to be errors, since they produce a result, albeit non-finite. Finally, the method for `call` evaluates the arguments to the `pow`, `sin`, or `sqrt` function, then calls the corresponding function in the math package.
```go
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
```
Several of these methods can fail. For example, a `call` expression could have an unknown function or the wrong number of arguments. It's also possible to construct a `unary` or `binary` expression with an invalid operator such as `!` or `<` (although the `Parse` function mentioned below will never do this). These errors cause `Eval` to panic. Other errors, like evaluating a `Var` not present in the environment, merely cause `Eval` to return the wrong result. All of these errors could be detected by inspecting the `Expr` before evaluating it. That will be the job of the `Check` method, which we will show soon, but first let's test `Eval`.

The `TestEval` function below is a test of the evaluator. It uses the `testing` package, which we'll explain in Chapter 11, but for now it's enough to know that calling `t.Errorf` reports an error. The function loops over a table of inputs that defines three expressions and different environments for each one. The first expression computes the radius of a circle given its area `A`, the second computes the sum of the cubes of two variables `x` and `y`, and the third converts a Fahrenheit temperature `F` to Celsius.
```go
// gopl.io/ch7/eval/eval_test
func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		//!-Eval
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		//!+Eval
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
```
For each entry in the table, the test parses the expression, evaluates it in the environment, and prints the result. We don't have space to show the `Parse` function here, but you'll find it if you download the package using `go get`.

The `go test` command (§11.1) runs a package's tests:
```bash
    $ go test -v gopl.io/ch7/eval
```
The `-v` flag lets us see the printed output of the test, which is normally suppressed for a suc- cessful test like this one. Here is the output of the test's `fmt.Printf` statements:
```
sqrt(A / pi)
	map[A:87616 pi:3.141592653589793] => 167

pow(x, 3) + pow(y, 3)
	map[x:12 y:1] => 1729
	map[x:9 y:10] => 1729

5 / 9 * (F - 32)
	map[F:-40] => -40
	map[F:32] => 0
	map[F:212] => 100
```
Fortunately the inputs so far have all been well formed, but our luck is unlikely to last. Even in interpreted languages, it is common to check the syntax for *static* errors, that is, mistakes that can be detected without running the program. By separating the static checks from the dynamic ones, we can detect errors sooner and perform many checks only once instead of each time an expression is evaluated.

Let's add another method to the `Expr` interface. The `Check` method checks for static errors in an expression syntax tree. We'll explain its `vars` parameter in a moment.
```go
    type Expr interface {
        Eval(env Env) float64
        // Check reports errors in this Expr and adds its Vars to the set.
        Check(vars map[Var]bool) error
    }
```
TheconcreteCheckmethodsareshownbelow. Evaluation of `literal` and `Var` cannot fail, so the `Check` methods for these types return `nil`. The methods for `unary` and `binary` first check that the operator is valid, then recursively check the operands. Similarly, the method for `call` first checks that the function is known and has the right number of arguments, then recursively checks each argument.
```go
// gopl.io/ch7/eval/check.go
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}
```
We've listed a selection of flawed inputs and the errors they elicit, in two groups. The `Parse` function (not shown) reports syntax errors and the `Check` function reports semantic errors.
```
x % 2         unexpected '%'
math.Pi       unexpected '.'
!true         unexpected '!'
"hello"       unexpected '"'
log(10)       unknown function "log"
sqrt(1, 2)    call to sqrt has 2 args, want 1
```
`Check`'s argument, a set of `Vars`, accumulates the set of variable names found within the expression. Each of these variables must be present in the environment for evaluation to succeed. This set is logically the result of the call to `Check`, but because the method is recursive, it is more convenient for `Check` to populate a set passed as a parameter. The client must provide an empty set in the initial call.

In Section 3.2, we plotted a function `f(x,y)` that was fixed at compile time. Now that we can parse, check, and evaluate expressions in strings, we can build a web application that receives an expression at run time from the client and plots the surface of that function. We can use the `vars` set to check that the expression is a function of only two variables, `x` and `y`—three, actually, since we'll provide `r`, the radius, as a convenience. And we'll use the `Check` method to reject ill-formed expressions before evaluation begins so that we don't repeat those checks during the 40,000 evaluations (100x100 cells, each with four corners) of the function that follow.

The `parseAndCheck` function combines these parsing and checking steps:
```go
// gopl.io/ch7/surface
import "gopl.io/ch7/eval"

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}
```
To make this a web application, all we need is the `plot` function below, which has the familiar signature of an `http.HandlerFunc`:
```go
// gopl.io/ch7/surface
func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0,0)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}
```

![Figure 7.7](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig7.7.png)

The `plot` function parses and checks the expression specified in the HTTP request and uses it to create an anonymous function of two variables. The anonymous function has the same signature as the fixed function `f` from the original surface-plotting program, but it evaluates the user-supplied expression. The environment defines `x`, `y`, and the radius `r`. Finally, `plot` calls `surface`, which is just the `main` function from `gopl.io/ch3/surface`, modified to take the function to plot and the output `io.Writer` as parameters, instead of using the fixed function `f` and `os.Stdout`. Figure 7.7 shows three surfaces produced by the program.

### Exercises
- **Exercise 7.13**: Add a `String` method to `Expr` to pretty-print the syntax tree. Check that the results, when parsed again, yield an equivalent tree.
- **Exercise 7.14**: Define a new concrete type that satisfies the `Expr` interface and provides a new operation such as computing the minimum value of its operands. Since the `Parse` function does not create instances of this new type, to use it you will need to construct a syntax tree directly (or extend the parser).
- **Exercise 7.15**: Write a program that reads a single expression from the standard input, prompts the user to provide values for any variables, then evaluates the expression in the resulting environment. Handle all errors gracefully.
- **Exercise 7.16**: Write a web-based calculator program.


## 7.10. Type Assertions 

A *type assertion* is an operation applied to an interface value. Syntactically, it looks like `x.(T)`, where `x` is an expression of an interface type and `T` is a type, called the "asserted" type. A type assertion checks that the dynamic type of its operand matches the asserted type.

There are two possibilities. First, if the asserted type `T` is a concrete type, then the type assertion checks whether `x`'s dynamic type is identical to `T`. If this check succeeds, the result of the type assertion is `x`'s dynamic value, whose type is of course `T`. In other words, a type assertion to a concrete type extracts the concrete value from its operand. If the check fails, then the operation panics. For example:
```go
    var w io.Writer
    w = os.Stdout
    f := w.(*os.File)      // success: f == os.Stdout
    c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
```
Second, if instead the asserted type `T` is an interface type, then the type assertion checks whether `x`'s dynamic type *satisfies* `T`. If this check succeeds, the dynamic value is not extracted; the result is still an interface value with the same type and value components, but the result has the interface type `T`. In other words, a type assertion to an interface type changes the type of the expression, making a different (and usually larger) set of methods accessible, but it preserves the dynamic type and value components inside the interface value.

After the first type assertion below, both `w` and `rw` hold `os.Stdout` so each has a dynamic type of `*os.File`, but `w`, an `io.Writer`, exposes only the file's `Write` method, whereas `rw` exposes its `Read` method too.
```go
    var w io.Writer

    w = os.Stdout
    rw := w.(io.ReadWriter) // success: *os.File has both Read and Write

    w = new(ByteCounter)
    rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
```
No matter what type was asserted, if the operand is a nil interface value, the type assertion fails. A type assertion to a less restrictive interface type (one with fewer methods) is rarely needed, as it behaves just like an assignment, except in the nil case.
```go
    w = rw             // io.ReadWriter is assignable to io.Writer
    w = rw.(io.Writer) // fails only if rw == nil
```
Often we're not sure of the dynamic type of an interface value, and we'd like to test whether it is some particular type. If the type assertion appears in an assignment in which two results are expected, such as the following declarations, the operation does not panic on failure but instead returns an additional second result, a boolean indicating success:
```go
    var w io.Writer = os.Stdout
    f, ok := w.(*os.File)      // success:  ok, f == os.Stdout
    b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
```
The second result is conventionally assigned to a variable named `ok`. If the operation failed, `ok` is false, and the first result is equal to the zero value of the asserted type, which in this example is a nil `*bytes.Buffer`.

The ok result is often immediately used to decide what to do next. The extended form of the if statement makes this quite compact:
```go
    if f, ok := w.(*os.File); ok {
        // ...use f...
    }
```
When the operand of a type assertion is a variable, rather than invent another name for the new local variable, you'll sometimes see the original name reused, shadowing the original, like this:
```go
    if w, ok := w.(*os.File); ok {
        // ...use w...
    }
```


## 7.11. Discriminating Errors with Type Assertions 

Consider the set of errors returned by file operations in the os package. I/O can fail for any number of reasons, but three kinds of failure often must be handled differently: file already exists (for create operations), file not found (for read operations), and permission denied. The `os` package provides these three helper functions to classify the failure indicated by a given error value:
```go
    package os
  
    func IsExist(err error) bool
    func IsNotExist(err error) bool
    func IsPermission(err error) bool
```
A naïve implementation of one of these predicates might check that the error message contains a certain substring,
```go
    func IsNotExist(err error) bool {
        // NOTE: not robust!
        return strings.Contains(err.Error(), "file does not exist")
    }
```
but because the logic for handling I/O errors can vary from one platform to another, this approach is not robust and the same failure may be reported with a variety of different error messages. Checking for substrings of error messages may be useful during testing to ensure that functions fail in the expected manner, but it's inadequate for production code.

A more reliable approach is to represent structured error values using a dedicated type. The `os` package defines a type called `PathError` to describe failures involving an operation on a file path, like `Open` or `Delete`, and a variant called `LinkError` to describe failures of operations involving two file paths, like `Symlink` and `Rename`. Here's `os.PathError`:
```go
    package os

    // PathError records an error and the operation and file path that caused it.
    type PathError struct {
        Op   string
        Path string
        Err  error
    }

    func (e *PathError) Error() string {
        return e.Op + " " + e.Path + ": " + e.Err.Error()
    }
```
Most clients are oblivious to `PathError` and deal with all errors in a uniform way by calling their `Error` methods. Although `PathError`'s `Error` method forms a message by simply concatenating the fields, `PathError`'s structure preserves the underlying components of the error. Clients that need to distinguish one kind of failure from another can use a type assertion to detect the specific type of the error; the specific type provides more detail than a simple string.
```go
    _, err := os.Open("/no/such/file")
    fmt.Println(err) // "open /no/such/file: No such file or directory"
    fmt.Printf("%#v\n", err)
    // Output:
    // &os.PathError{Op:"open", Path:"/no/such/file", Err:0x2}
```
That's how the three helper functions work. For example, `IsNotExist`, shown below, reports whether an error is equal to `syscall.ENOENT` (§7.8) or to the distinguished error `os.ErrNotExist` (see `io.EOF` in §5.4.2), or is a `*PathError` whose underlying error is one of those two.
```go
    import (
        "errors"
        "syscall"
    )

    var ErrNotExist = errors.New("file does not exist")
    // IsNotExist returns a boolean indicating whether the error is known to
    // report that a file or directory does not exist. It is satisfied by
    // ErrNotExist as well as some syscall errors.
    func IsNotExist(err error) bool {
        if pe, ok := err.(*PathError); ok {
            err = pe.Err
        }
        return err == syscall.ENOENT || err == ErrNotExist
    }
```
And here it is in action:
```go
    _, err := os.Open("/no/such/file")
    fmt.Println(os.IsNotExist(err)) // "true"
```
Of course, `PathError`'s structure is lost if the error message is combined into a larger string, for instance by a call to `fmt.Errorf`. Error discrimination must usually be done immediately after the failing operation, before an error is propagated to the caller.


## 7.12. Querying Behaviors with Interface Type Assertions 

The logic below is similar to the part of the `net/http` web server responsible for writing HTTP header fields such as `"Content-type:text/html"`. The `io.Writer` `w` represents the HTTP response; the bytes written to it are ultimately sent to someone's web browser.
```go
    func writeHeader(w io.Writer, contentType string) error {
        if _, err := w.Write([]byte("Content-Type: ")); err != nil {
            return err
        }
        if _, err := w.Write([]byte(contentType)); err != nil {
            return err
        }
        // ...
    }
```
Because the `Write` method requires a byte slice, and the value we wish to write is a string, a `[]byte(...)` conversion is required. This conversion allocates memory and makes a copy, but the copy is thrown away almost immediately after. Let's pretend that this is a core part of the web server and that our profiling has revealed that this memory allocation is slowing it down. Can we avoid allocating memory here?

The `io.Writer` interface tells us only one fact about the concrete type that `w` holds: that bytes may be written to it. If we look behind the curtains of the `net/http` package, we see that the dynamic type that `w` holds in this program also has a `WriteString` method that allows strings to be efficiently written to it, avoiding the need to allocate a temporary copy. (This may seem like a shot in the dark, but a number of important types that satisfy `io.Writer` also have a `WriteString` method, including `*bytes.Buffer`, `*os.File`, and `*bufio.Writer`.)

We cannot assume that an arbitrary `io.Writer` `w` also has the `WriteString` method. But we can define a new interface that has just this method and use a type assertion to test whether the dynamic type of `w` satisfies this new interface.
```go
    // writeString writes s to w.
    // If w has a WriteString method, it is invoked instead of w.Write.
    func writeString(w io.Writer, s string) (n int, err error) {
        type stringWriter interface {
            WriteString(string) (n int, err error)
        }
        if sw, ok := w.(stringWriter); ok {
            return sw.WriteString(s) // avoid a copy
        }
        return w.Write([]byte(s)) // allocate temporary copy
    }

    func writeHeader(w io.Writer, contentType string) error {
        if _, err := writeString(w, "Content-Type: "); err != nil {
            return err
        }
        if _, err := writeString(w, contentType); err != nil {
            return err
        }
        // ...
    }
```
To avoid repeating ourselves, we've moved the check into the utility function `writeString`, but it is so useful that the standard library provides it as `io.WriteString`. It is the recommended way to write a string to an `io.Writer`.

What's curious in this example is that there is no standard interface that defines the `WriteString` method and specifies its required behavior. Furthermore, whether or not a concrete type satisfies the `stringWriter` interface is determined only by its methods, not by any declared relationship between it and the interface type. What this means is that the technique above relies on the assumption that *if* a type satisfies the interface below, *then* `WriteString(s)` must have the same effect as `Write([]byte(s))`.
```go
    interface {
        io.Writer
        WriteString(s string) (n int, err error)
    }
```
Although `io.WriteString` documents its assumption, few functions that call it are likely to document that they too make the same assumption. Defining a method of a particular type is taken as an implicit assent for a certain behavioral contract. Newcomers to Go, especially those from a background in strongly typed languages, may find this lack of explicit intention unsettling, but it is rarely a problem in practice. With the exception of the empty interface `interface{}`, interface types are seldom satisfied by unintended coincidence.

The writeString function above uses a type assertion to see whether a value of a general interface type also satisfies a more specific interface type, and if so, it uses the behaviors of the specific interface. This technique can be put to good use whether or not the queried interface is standard like `io.ReadWriter` or user-defined like stringWriter.

It's also how `fmt.Fprintf` distinguishes values that satisfy error or `fmt.Stringer` from all other values. Within `fmt.Fprintf`, there is a step that converts a single operand to a string, something like this:
```go
    package fmt

    func formatOneValue(x interface{}) string {
        if err, ok := x.(error); ok {
            return err.Error()
        }

        if str, ok := x.(Stringer); ok {
            return str.String()
        }
        // ...all other types...
    }
```
If `x` satisfies either of the two interfaces, that determines the formatting of the value. If not, the default case handles all other types more or less uniformly using reflection; we'll find out how in Chapter 12.

Again, this makes the assumption that any type with a `String` method satisfies the behavioral contract of `fmt.Stringer`, which is to return a string suitable for printing.


## 7.13. Type Switches 

Interfaces are used in two distinct styles. In the first style, exemplified by `io.Reader`, `io.Writer`, `fmt.Stringer`, `sort.Interface`, `http.Handler`, and `error`, an interface's methods express the similarities of the concrete types that satisfy the interface but hide the representation details and intrinsic operations of those concrete types. The emphasis is on the methods, not on the concrete types.

The second style exploits the ability of an interface value to hold values of a variety of concrete types and considers the interface to be the *union* of those types. Type assertions are used to discriminate among these types dynamically and treat each case differently. In this style, the emphasis is on the concrete types that satisfy the interface, not on the interface's methods (if indeed it has any), and there is no hiding of information. We'll describe interfaces used this way as *discriminated unions*.

If you're familiar with object-oriented programming, you may recognize these two styles as *subtype polymorphism* and *ad hoc polymorphism*, but you needn't remember those terms. For the remainder of this chapter, we'll present examples of the second style.

Go's API for querying an SQL database, like those of other languages, lets us cleanly separate the fixed part of a query from the variable parts. An example client might look like this:
```go
    import "database/sql"

    func listTracks(db sql.DB, artist string, minYear, maxYear int) {
        result, err := db.Exec(
            "SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
            artist, minYear, maxYear)
        // ...
    }
```

The `Exec` method replaces each `'?'` in the query string with an SQL literal denoting the corresponding argument value, which may be a boolean, a number, a string, or `nil`. Constructing queries this way helps avoid SQL injection attacks, in which an adversary takes control of the query by exploiting improper quotation of input data. Within `Exec`, we might find a function like the one below, which converts each argument value to its literal SQL notation.
```go
    func sqlQuote(x interface{}) string {
        if x == nil {
            return "NULL"
        } else if _, ok := x.(int); ok {
            return fmt.Sprintf("%d", x)
        } else if _, ok := x.(uint); ok {
            return fmt.Sprintf("%d", x)
        } else if b, ok := x.(bool); ok {
            if b {
                return "TRUE"
            }
            return "FALSE"
        } else if s, ok := x.(string); ok {
            return sqlQuoteString(s) // (not shown)
        } else {
            panic(fmt.Sprintf("unexpected type %T: %v", x, x))
        } 
    }
```
A `switch` statement simplifies an `if`-`else` chain that performs a series of value equality tests. An analogous *type switch* statement simplifies an `if`-`else` chain of type assertions.

In its simplest form, a type switch looks like an ordinary switch statement in which the operand is `x.(type)` (that's literally the keyword `type`) and each case has one or more types. A type switch enables a multi-way branch based on the interface value's dynamic type. The nil case matches if `x == nil`, and the default case matches if no other case does. A type switch for `sqlQuote` would have these cases:
```go
    switch x.(type) {
        case nil:       // ...
        case int, uint: // ...
        case bool:      // ...
        case string:    // ...
        default:        // ...
    }
```
As with an ordinary `switch` statement (§1.8), cases are considered in order and, when a match is found, the case's body is executed. Case order becomes significant when one or more case types are interfaces, since then there is a possibility of two cases matching. The position of the default case relative to the others is immaterial. No fallthrough is allowed.

Notice that in the original function, the logic for the `bool` and `string` cases needs access to the value extracted by the type assertion. Since this is typical, the type switch statement has an extended form that binds the extracted value to a new variable within each case:
```go
    switch x := x.(type) { /* ... */ }
```
Here we've called the new variables `x` too; as with type assertions, reuse of variable names is common. Like a `switch` statement, a type switch implicitly creates a lexical block, so the declaration of the new variable called `x` does not conflict with a variable `x` in an outer block. Each case also implicitly creates a separate lexical block.

Rewriting `sqlQuote` to use the extended form of type switch makes it significantly clearer:
```go
func sqlQuote(x interface{}) string {
         switch x := x.(type) {
         case nil:
             return "NULL"
         case int, uint:
             return fmt.Sprintf("%d", x) // x has type interface{} here.
         case bool:
             if x {
                 return "TRUE"
}
             return "FALSE"
         case string:
             return sqlQuoteString(x) // (not shown)
         default:
} }
panic(fmt.Sprintf("unexpected type %T: %v", x, x))
```
In this version, within the block of each single-type case, the variable x has the same type as thecase. For instance, `x` has type `bool` within the `bool` case and `string` within the `string` case. In all other cases, `x` has the (interface) type of the `switch` operand, which is `interface{}` in this example. When the same action is required for multiple cases, like `int` and `uint`, the type switch makes it easy to combine them.

Although `sqlQuote` accepts an argument of any type, the function runs to completion only if the argument's type matches one of the cases in the type switch; otherwise it panics with an "unexpected type" message. Although the type of `x` is `interface{}`, we consider it a discriminated union of `int`, `uint`, `bool`, `string`, and `nil`.


## 7.14. Example: Token-Based XML Decoding 

Section 4.5 showed how to decode JSON documents into Go data structures with the `Marshal` and `Unmarshal` functions from the `encoding/json` package. The `encoding/xml` package provides a similar API. This approach is convenient when we want to construct a representation of the document tree, but that's unnecessary for many programs. The `encoding/xml` package also provides a lower-level token-based API for decoding XML. In the token-based style, the parser consumes the input and produces a stream of tokens, primarily of four kinds (`StartElement`, `EndElement`, `CharData`, and `Comment`) each being a concrete type in the `encoding/xmlpackage`. Each call to `(*xml.Decoder)`. `Token` returns a token. 

The relevant parts of the API are shown here:
```go
// encoding/xml
    package xml

    type Name struct {
        Local string // e.g., "Title" or "id"
    }
    type Attr struct { // e.g., name="value"
        Name  Name
        Value string
    }

    // A Token includes StartElement, EndElement, CharData,
    // and Comment, plus a few esoteric types (not shown).
    type Token interface{}
    type StartElement struct { // e.g., <name>
        Name Name
        Attr []Attr
    }
    type EndElement struct { Name Name } // e.g., </name>
    type CharData []byte                 // e.g., <p>CharData</p>
    type Comment []byte                  // e.g., <!-- Comment -->

    type Decoder struct{ /* ... */ }

    func NewDecoder(io.Reader) *Decoder
    func (*Decoder) Token() (Token, error) // returns next Token in sequence
```
The `Token` interface, which has no methods, is also an example of a discriminated union. The purpose of a traditional interface like `io.Reader` is to hide details of the concrete types that satisfy it so that new implementations can be created; each concrete type is treated uniformly. By contrast, the set of concrete types that satisfy a discriminated union is fixed by the design and exposed, not hidden. Discriminated union types have few methods; functions that operate on them are expressed as a set of cases using a type switch, with different logic in each case.

The `xmlselect` program below extracts and prints the text found beneath certain elements in an XML document tree. Using the API above, it can do its job in a single pass over the input without ever materializing the tree.
```go
// gopl.io/ch7/xmlselect
// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
```
Each time the loop in `main` encounters a `StartElement`, it pushes the element's name onto a stack, and for each `EndElement` it pops the name from the stack. The API guarantees that the sequence of `StartElement` and `EndElement` tokens will be properly matched, even in ill-formed documents. `Comments` are ignored. When xmlselect encounters a `CharData`, it prints the text only if the stack contains all the elements named by the command-line arguments, in order.

The command below prints the text of any h2 elements appearing beneath two levels of `div` elements. Its input is the XML specification, itself an XML document.
```
    $ go build gopl.io/ch1/fetch
    $ ./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 |
        ./xmlselect div div h2
    html body div div h2: 1 Introduction
    html body div div h2: 2 Documents
    html body div div h2: 3 Logical Structures
    html body div div h2: 4 Physical Structures
    html body div div h2: 5 Conformance
    html body div div h2: 6 Notation
    html body div div h2: A References
    html body div div h2: B Definitions for Character Normalization
    ...
```

### Exercises
- **Exercise 7.17**: Extend `xmlselect` so that elements may be selected not just by name, but by their attributes too, in the manner of CSS, so that, for instance, an element like `<div id="page" class="wide">` could be selected by a matching `id` or `class` as well as its name.
- **Exercise 7.18**: Using the token-based decoder API, write a program that will read an arbitrary XML document and construct a tree of generic nodes that represents it. Nodes are of two kinds: `CharData` nodes represent text strings, and `Element` nodes represent named elements and their attributes. Each element node has a slice of child nodes.

You may find the following declarations helpful.
```go
    import "encoding/xml"

    type Node interface{} // CharData or *Element
     
    type CharData string

    type Element struct {
        Type     xml.Name
        Attr     []xml.Attr
        Children []Node
    }
```





## 7.15. A Few Words of Advice

When designing a new package, novice Go programmers often start by creating a set of interfaces and only later define the concrete types that satisfy them. This approach results in many interfaces, each of which has only a single implementation. Don't do that. Such interfaces are unnecessary abstractions; they also have a run-time cost. You can restrict which methods of a type or fields of a struct are visible outside a package using the export mechanism (§6.6). Interfaces are only needed when there are two or more concrete types that must be dealt with in a uniform way.

We make an exception to this rule when an interface is satisfied by a single concrete type but that type cannot live in the same package as the interface because of its dependencies. In that case, an interface is a good way to decouple two packages.

Because interfaces are used in Go only when they are satisfied by two or more types, they necessarily abstract away from the details of any particular implementation. The result is smaller interfaces with fewer, simpler methods, often just one as with `io.Writer` or `fmt.Stringer`. Small interfaces are easier to satisfy when new types come along. A good rule of thumb for interface design is *ask only for what you need*.

This concludes our tour of methods and interfaces. Go has great support for the object-oriented style of programming, but this does not mean you need to use it exclusively. Not everything need be an object; standalone functions have their place, as do unencapsulated data types. Observe that together, the examples in the first five chapters of this book call no more than two dozen methods, like `input.Scan`, as opposed to ordinary function calls like `fmt.Printf`.
