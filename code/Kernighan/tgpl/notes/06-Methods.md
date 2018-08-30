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

  perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
  fmt.Println(geometry.PathDistance(perim)) // "12", standalone function
  fmt.Println(perim.Distance())             // "12", method of geometry.Pat
```


## 6.2. Methods with a Pointer Receiver 


## 6.3. Composing Types by Struct Embedding 
## 6.4. Method Values and Expressions 
## 6.5. Example: Bit Vector Type 
## 6.6. Encapsulation
