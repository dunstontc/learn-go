# Chapter 6: Methods

<!-- TOC -->

- [6.1. Method Declarations](#61-method-declarations)
- [6.2. Methods with a Pointer Receiver](#62-methods-with-a-pointer-receiver)
- [6.3. Composing Types by Struct Embedding](#63-composing-types-by-struct-embedding)
- [6.4. Method Values and Expressions](#64-method-values-and-expressions)
- [6.5. Example: Bit Vector Type](#65-example-bit-vector-type)
- [6.6. Encapsulation](#66-encapsulation)

<!-- /TOC -->

Since the early 1990s, object-oriented programming (OOP) has been the dominant programming paradigm in industry and education, and nearly all widely used languages developed since then have included support for it. Go is no exception.

Although there is no universally accepted definition of object-oriented programming, for our purposes, an object is simply a value or variable that has methods, and a method is a function associated with a particular type. An object-oriented program is one that uses methods to express the properties and operations of each data structure so that clients need not access the object’s representation directly.

In earlier chapters, we have made regular use of methods from the standard library, like the Seconds method of type time.Duration:
```go
  const day = 24 * time.Hour
  fmt.Println(day.Seconds()) // "86400"
```
and we defined a method of our own in Section 2.5, a String method for the `Celsius` type:
```go
  func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```
In this chapter, the first of two on object-oriented programming, we’ll show how to define and use methods effectively. We’ll also cover two key principles of object-oriented programming, *encapsulation* and *composition*.


## 6.1. Method Declarations 

A method is declared with a variant of the ordinary function declaration in which an extra parameter appears before the function name. The parameter attaches the function to the type of that parameter.

Let’s write our first method in a simple package for plane geometry:
```go
// gopl.io/ch6/geometry
// Package geometry defines simple types for plane geometry.
package geometry

import "math"

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot((q.X - p.X), (q.Y - p.Y))
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot((q.X-p.X), ((q.Y-p.Y))
}
```
The extra parameter p is called the method’s *receiver*, a legacy from early object-oriented languages that described calling a method as "sending a message to an object."

In Go, we don’t use a special name like *this* or *self* for the receiver; we choose receiver names just as we would for any other parameter. Since the receiver name will be frequently used, it’s a good idea to choose something short and to be consistent across methods. A common choice is the first letter of the type name, like `p` for `Point`.

In a method call, the receiver argument appears before the method name. This parallels the declaration, in which the receiver parameter appears before the method name.
```go
  p := Point{1, 2}
  q := Point{4, 6}
  fmt.Println(Distance(p, q)) // "5", function call
  fmt.Println(p.Distance(q))  // "5", method call
```
There’s no conflict between the two declarations of functions called `Distance` above. The first declares a package-level function called `geometry.Distance`. The second declares a method of the type `Point`, so its name is `Point.Distance`.

The expression `p.Distance` is called a *selector*, because it selects the appropriate `Distance` method for the receiver `p` of type `Point`. Selectors are also used to select fields of struct types, as in `p.X`. Since methods and fields inhabit the same name space, declaring a method `X` on the struct type `Point` would be ambiguous and the compiler will reject it.

Since each type has its own name space for methods, we can use the name `Distance` for other methods so long as they belong to different types. Let’s define a type `Path` that represents a sequence of line segments and give it a `Distance` method too.
```go
  // A Path is a journey connecting the points with straight lines.
  type Path []Point

  // Distance returns the distance traveled along the path.
  func (path Path) Distance() float64 {
      sum := 0.0
      for i := range path {
          if i > 0 {
              sum += path[i-1].Distance(path[i])
          } 
      }
      return sum
  }
```
`Path` is a named slice type, not a struct type like `Point`, yet we can still define methods for it. In allowing methods to be associated with any type, Go is unlike many other object-oriented languages. It is often convenient to define additional behaviors for simple types such as numbers, strings, slices, maps, and sometimes even functions. Methods may be declared on any named type defined in the same package, so long as its underlying type is neither a pointer nor an interface.

The two `Distance` methods have different types. They’re not related to each other at all, though `Path.Distance` uses `Point.Distance` internally to compute the length of each segment that connects adjacent points.

Let’s call the new method to compute the perimeter of a right triangle:
```go
  perim := Path{
      {1, 1},
      {5, 1},
      {5, 4},
      {1, 1},
  }
  fmt.Println(perim.Distance()) // "12"
```
In the two calls above to methods named `Distance`, the compiler determines which function to call based on both the method name and the type of the receiver. In the first, `path[i-1]` has type `Point` so `Point.Distance` is called; in the second, perim has type `Path`, so `Path.Distance` is called.

All methods of a given type must have unique names, but different types can use the same name for a method, like the `Distance` methods for `Point` and `Path`; there’s no need to qualify function names (for example, `PathDistance`) to disambiguate. Here we see the first benefit to using methods over ordinary functions: method names can be shorter. The benefit is magnified for calls originating outside the package, since they can use the shorter name *and* omit the package name:
```go
  import "gopl.io/ch6/geometry"

  perim := geometry.Path{
      {1, 1}, 
      {5, 1}, 
      {5, 4}, 
      {1, 1},
  }

  fmt.Println(geometry.PathDistance(perim)) // "12", standalone function
  fmt.Println(perim.Distance())             // "12", method of geometry.Pat
```


## 6.2. Methods with a Pointer Receiver 

Because calling a function makes a copy of each argument value, if a function needs to update a variable, or if an argument is so large that we wish to avoid copying it, we must pass the address of the variable using a pointer. The same goes for methods that need to update the receiver variable: we attach them to the pointer type, such as `*Point`.

```go
  func (p *Point) ScaleBy(factor float64) {
      p.X *= factor
      p.Y *= factor
  }
```
The name of this method is `(*Point).ScaleBy`. The parentheses are necessary; without them, the expression would be parsed as `*(Point.ScaleBy)`.

In a realistic program, convention dictates that if any method of Point has a pointer receiver, then all methods of Point should have a pointer receiver, even ones that don’t strictly need it. We’ve broken this rule for `Point` so that we can show both kinds of method.

Named types (`Point`) and pointers to them (`*Point`) are the only types that may appear in a receiver declaration. Furthermore, to avoid ambiguities, method declarations are not permitted on named types that are themselves pointer types:
```go
  type P *int
  func (P) f() { /* ... */ } // compile error: invalid receiver type
```
The `(*Point).ScaleBy` method can be called by providing a `*Point` receiver, like this:
```go
  r := &Point{1, 2}
  r.ScaleBy(2)
  fmt.Println(*r) // "{2, 4}"
```
or this:
```go
  p := Point{1, 2}
  pptr := &p
  pptr.ScaleBy(2)
  fmt.Println(p) // "{2, 4}"
```
or this:
```go
  p := Point{1, 2}
  (&p).ScaleBy(2)
  fmt.Println(p) // "{2, 4}"
```
But the last two cases are ungainly. Fortunately, the language helps us here. If the receiver `p` is a *variable* of type `Point` but the method requires a `*Point` receiver, we can use this shorthand:
```go
  p.ScaleBy(2)
```
and the compiler will perform an implicit `&p` on the variable. This works only for variables, including struct fields like `p.X` and array or slice elements like `perim[0]`. We cannot call a `*Point` method on a non-addressable `Point` receiver, because there’s no way to obtain the address of a temporary value.
```go
  Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal
```
But we can call a `Point` method like `Point.Distance` with a `*Point` receiver, because there is a way to obtain the value from the address: just load the value pointed to by the receiver. The compiler inserts an implicit `*` operation for us. These two function calls are equivalent:
```go
  pptr.Distance(q)
  (*pptr).Distance(q)
```
Let’s summarize these three cases again, since they are a frequent point of confusion. In every valid method call expression, exactly one of these three statements is true.

Either the receiver argument has the same type as the receiver parameter, for example both have type `T` or both have type `*T`:
```go
  Point{1, 2}.Distance(q) //  Point
  pptr.ScaleBy(2)         // *Point
```
Or the receiver argument is a variable of type `T` and the receiver parameter has type `*T`. The compiler implicitly takes the address of the variable:
```go
  p.ScaleBy(2) // implicit (&p)
```
Or the receiver argument has type `*T` and the receiver parameter has type `T`. The compiler implicitly dereferences the receiver, in other words, loads the value:
```go
  pptr.Distance(q) // implicit (*pptr)
```
If all the methods of a named type `T` have a receiver type of `T` itself (not `*T`), it is safe to copy instances of that type; calling any of its methods necessarily makes a copy. For example, `time.Duration` values are liberally copied, including as arguments to functions. But if any method has a pointer receiver, you should avoid copying instances of `T` because doing so may violate internal invariants. For example, copying an instance of `bytes.Buffer` would cause the original and the copy to alias (§2.3.2) the same underlying array of bytes. Subsequent method calls would have unpredictable effects.

### 6.2.1 Nil Is a Valid Receiver Value

Just as some functions allow nil pointers as arguments, so do some methods for their receiver, especially if `nil` is a meaningful zero value of the type, as with maps and slices. In this simple linked list of integers, `nil` represents the empty list:
```go
  // An IntList is a linked list of integers.
  // A nil *IntList represents the empty list.
  type IntList struct {
      Value int
      Tail  *IntList
  }
  
  // Sum returns the sum of the list elements.
  func (list *IntList) Sum() int {
      if list == nil {
          return 0
      }
      return list.Value + list.Tail.Sum()
  }
```
When you define a type whose methods allow nil as a receiver value, it’s worth pointing this
out explicitly in its documentation comment, as we did above.

Here’s part of the definition of the `Values` type from the `net/url` package:
```go
  // net/url
  package url

  // Values maps a string key to a list of values.
  type Values map[string][]string

  // Get returns the first value associated with the given key,
  // or "" if there are none.
  func (v Values) Get(key string) string {
      if vs := v[key]; len(vs) > 0 {
          return vs[0]
      }
      return ""
  }

  // Add adds the value to key.
  // It appends to any existing values associated with key.
  func (v Values) Add(key, value string) {
      v[key] = append(v[key], value)
  }
```
It exposes its representation as a map but also provides methods to simplify access to the map, whose values are slices of strings—it’s a *multimap*. Its clients can use its intrinsic operators (`make`, slice literals, `m[key]`, and so on), or its methods, or both, as they prefer:
```go
// gopl.io/ch6/urlvalues

```
In the final call to `Get`, the nil receiver behaves like an empty map. We could equivalently have written it as `Values(nil).Get("item"))`, but `nil.Get("item")` will not compile because the type of `nil` has not been determined. By contrast, the final call to `Add` panics as it tries to update a nil map.

Because `url.Values` is a map type and a map refers to its key/value pairs indirectly, any updates and deletions that `url.Values.Add` makes to the map elements are visible to the caller. However, as with ordinary functions, any changes a method makes to the reference itself, like setting it to n il or making it refer to a different map data structure, will not be reflected in the caller.


## 6.3. Composing Types by Struct Embedding 

Consider the type `ColoredPoint`:


## 6.4. Method Values and Expressions 
## 6.5. Example: Bit Vector Type 
## 6.6. Encapsulation
