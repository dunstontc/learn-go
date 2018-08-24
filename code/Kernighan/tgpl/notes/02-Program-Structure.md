# Chapter 2: Program Structure

<!-- TOC -->

- [2.1. Names](#21-names)
- [2.2. Declarations](#22-declarations)
- [2.3. Variables](#23-variables)
  - [2.3.1 Short Variable Declarations](#231-short-variable-declarations)
  - [2.3.2 Pointers](#232-pointers)
  - [2.3.3 The `new` Function](#233-the-new-function)
  - [2.3.4 Lifetime of Variables](#234-lifetime-of-variables)
  - [2.3.5 Tuple Assignment](#235-tuple-assignment)
  - [2.3.6 Assignability](#236-assignability)
- [2.4. Assignments](#24-assignments)
- [2.5. Type Declarations](#25-type-declarations)
- [2.6. Packages and Files](#26-packages-and-files)
- [2.7. Scope](#27-scope)

<!-- /TOC -->

In Go, as in any other programming language, one builds large programs from a small set of basic constructs. Variables store values. Simple expressions are combined into larger ones with operations like addition and subtraction. Basic types are collected into aggregates like arrays and structs. Expressions are used in statements whose execution order is determined by control-flow statements like if and for. Statements are grouped into functions for isolation and reuse. Functions are gathered into source files and packages.  

We saw examples of most of these in the previous chapter. In this chapter, we’ll go into more detail about the basic structural elements of a Go program. The example programs are intentionally simple, so we can focus on the language without getting sidetracked by complicated algorithms or data structures.  


## 2.1. Names

The names of Go functions, variables, constants, types, statement labels, and packages follow a simple rule: a name begins with a letter (that is, anything that Unicode deems a letter) or an underscore and may have any number of additional letters, digits, and underscores. Case matters: `heapSort` and `Heapsort` are different names.  

Go has 25 [*keywords*](https://golang.org/ref/spec#Keywords) like `if` and `switch` that may be used only where the syntax permits; they can’t be used as names.  
```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```


In addition, there are about three dozen *predeclared* names like int and true for built-in constants, types, and functions:

Constants:
```
true false iota nil
```

Types:
```
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex64 compled128
bool byte rune string error
```

[Functions](https://golang.org/pkg/builtin/):
  - `make`
  - `len`
  - `cap`
  - `new`
  - `append`
  - `copy`
  - `close`
  - `delete`
  - `complex`
  - `real`
  - `imag`
  - `panic`
  - `recover`

These names are not reserved, so you may use them in declarations. We’ll see a handful of places where redeclaring one of them makes sense, but beware of the potential for confusion.  

If an entity is declared within a function, it is *local* to that function. If declared outside of a function, however, it is visible in all files of the package to which it belongs. The case of the first letter of a name determines its visibility across package boundaries. If the name begins with an upper-case letter, it is *exported*, which means that it is visible and accessible outside of its own package and may be referred to by other parts of the program, as with Printf in the fmt package. Package names themselves are always in lower case.  

There is no limit on name length, but convention and style in Go programs lean toward short names, especially for local variables with small scopes; you are much more likely to see variables named `i` than `theLoopIndex`. Generally, the larger the scope of a name, the longer and more meaningful it should be.  

Stylistically, Go programmers use *"camel case"* when forming names by combining words; that is, interior capital letters are preferred over interior underscores. Thus the standard libraries have functions with names like `QuoteRuneToASCII` and `parseRequestLine` but never `quote_rune_to_ASCII` or `parse_request_line`. The letters of acronyms and initialisms like ASCII and HTML are always rendered in the same case, so a function might be called `htmlEscape`, `HTMLEscape`, or `escapeHTML`, but not `escapeHtml`.  


## 2.2. Declarations

A *declaration* names a program entity and specifies some or all of its properties. There are four major kinds of declarations: var, const, type, and func. We’ll talk about variables and types in this chapter, constants in Chapter 3, and functions in Chapter 5.

A Go program is stored in one or more files whose names end in `.go`. Each file begins with a `package` declaration that says what package the file is part of. The `package` declaration is followed by any `import` declarations, and then a sequence of *package-level* declarations of types, variables, constants, and functions, in any order. For example, this program declares a constant, a function, and a couple of variables:  
```go
// gopl.io/ch2/boiling
// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}
```

The constant `boilingF` is a package-level declaration (as is `main`), whereas the variables `f` and `c` are local to the function `main`. The name of each package-level entity is visible not only throughout the source file that contains its declaration, but throughout all the files of the package. By contrast, local declarations are visible only within the function in which they are declared and perhaps only within a small part of it.  

<!-- A function declaration has a name, a list of parameters (the variables whose values are provided by the function’s callers), an optional list of results, and the function body, which contains the statements that define what the function does. The result list is omitted if the function does not return anything. Execution of the function begins with the first statement and continues until it encounters a return statement or reaches the end of a function that has no results. Control and any results are then returned to the caller.   -->
A function declaration has:
- a name 
- a list of *parameters* (the variables whose values are provided by the function’s callers) 
- an optional list of results
- the function *body* (which contains the statements that define what the function does)

The result list is omitted if the function does not return anything. Execution of the function begins with the first statement and continues until it encounters a return statement or reaches the end of a function that has no results. Control and any results are then returned to the caller.  

We’ve seen a fair number of functions already and there are lots more to come, including an extensive discussion in Chapter 5, so this is only a sketch. The function `fToC` below encapsulates the temperature conversion logic so that it is defined only once but may be used from multiple places. Here `main` calls it twice, using the values of two different local constants:  
```go
// gopl.io/ch2/ftoc
// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
```


## 2.3. Variables

A var declaration creates a variable of a particular type, attaches a name to it, and sets its initial value.  
Each declaration has the general form
```go
  var name type = expression
```
Either the type or the `= expression` part may be omitted, but not both. If the type is omitted, it is determined by the initializer expression. If the expression is omitted, the initial value is the *zero value* for the type, which is `0` for numbers, `false` for booleans, `""` for strings, and `nil` for interfaces and reference types (slice, pointer, map, channel, function). The zero value of an aggregate type like an array or a struct has the zero value of all of its elements or fields.  

The zero-value mechanism ensures that a variable always holds a well-defined value of its type; in Go there is no such thing as an uninitialized variable. This simplifies code and often ensures sensible behavior of boundary conditions without extra work.  
For example,  
```go
  var s string
  fmt.Println(s) // ""
```
prints an empty string, rather than causing some kind of error or unpredictable behavior. Go programmers often go to some effort to make the zero value of a more complicated type meaningful, so that variables begin life in a useful state.  

It is possible to declare and optionally initialize a set of variables in a single declaration, with a matching list of expressions. Omitting the type allows declaration of multiple variables of different types:  
```go
  var i, j, k int                 // int, int, int
  var b, f, s = true, 2.3, "four" // bool, float64, string
```
Initializers may be literal values or arbitrary expressions. Package-level variables are initialized before `main` begins (§2.6.2), and local variables are initialized as their declarations are encountered during function execution.   

A set of variables can also be initialized by calling a function that returns multiple values:  
```go
  var f, err = os.Open(name) // os.Open returns a file and an error
```

### 2.3.1 Short Variable Declarations
Within a function, an alternate form called a *short variable declaration* may be used to declare and initialize local variables. It takes the form `name := expression`, and the type of `name` is determined by the type of expression. Here are three of the many short variable declarations in the `lissajous` function (§1.4):  
```go
  anim := gif.GIF{LoopCount: nframes}
  freq := rand.Float64() * 3.0
  t := 0.0
```

Because of their brevity and flexibility, short variable declarations are used to declare and initialize the majority of local variables. A `var` declaration tends to be reserved for local variables that need an explicit type that differs from that of the initializer expression, or for when the variable will be assigned a value later and its initial value is unimportant.  
```go
  i := 100                  // an int
  var boiling float64 = 100 // a float64
  var names []string
  var err error
  var p Point
```

As with var declarations, multiple variables may be declared and initialized in the same short variable declaration,  
```go
  i, j := 0, 1
```
but declarations with multiple initializer expressions should be used only when they help readability, such as for short and natural groupings like the initialization part of a for loop.  

Keep in mind that `:=` is a declaration, whereas `=` is an assignment. A multi-variable declaration should not be confused with a *tuple assignment* (§2.4.1), in which each variable on the left-hand side is assigned the corresponding value from the right-hand side:
```go
  i, j = j, i // swap values of i and j
```

Like ordinary `var` declarations, short variable declarations may be used for calls to functions like `os.Open` that return two or more values:
```go
  f, err := os.Open(name)
    if err != nil {
  return err
  }
  // ...use f...
  f.Close()
```

One subtle but important point: a short variable declaration does not necessarily *declare* all the variables on its left-hand side. If some of them were already declared in the same lexical block (§2.7), then the short variable declaration acts like an *assignment* to those variables.  
In the code below, the first statement declares both `in` and `err`. The second declares `out` but only assigns a value to the existing `err` variable.  
```go
  in, err := os.Open(infile)
  // ...
  out, err := os.Create(outfile)
```
A short variable declaration must declare at least one new variable, however, so this code will not compile:
```go

  f, err := os.Open(infile)
  // ...
  f, err := os.Create(outfile) // compile error: no new variables
```
The fix is to use an ordinary assignment for the second statement.  

A short variable declaration acts like an assignment only to variables that were already declared in the same lexical block; declarations in an outer block are ignored. We’ll see examples of this at the end of the chapter.

### 2.3.2 Pointers

A variable is a piece of storage containing a value. Variables created by declarations are identified by a name, such as x, but many variables are identified only by expressions like x[i] or x.f. All these expressions read the value of a variable, except when they appear on the lefthand side of an assignment, in which case a new value is assigned to the variable.  

A *pointer* value is the *address* of a variable. A pointer is thus the location at which a value is stored. Not every value has an address, but every variable does. With a pointer, we can read or update the value of a variable *indirectly*, without using or even knowing the name of the variable, if indeed it has a name.  

If a variable is declared `var x int`, the expression `&x` ("address of x") yields a pointer to an integer variable, that is, a value of type `*int`, which is pronounced "pointer to int". If this value is called `p`, we say "p points to x", or equivalently "p contains the address of x". The variable to which `p` points is written `*p`. The expression `*p` yields the value of that variable, an `int`, but since `*p` denotes a variable, it may also appear on the left-hand side of an assignment, in which case the assignment updates the variable.  
```go
  x := 1
  p := &x         // p, of type *int, points to x
  fmt.Println(*p) // "1"
  *p = 2          // equivalent to x = 2
  fmt.Println(x)  // "2"
```

Each component of a variable of aggregate type (a field of a struct or an element of an array) is also a variable and thus has an address too. 

Variables are sometimes described as *addressable* values. Expressions that denote variables are the only expressions to which the *address-of* operator & may be applied.

The zero value for a pointer of any type is `nil`. The test `p != nil` is true if `p` points to a variable. Pointers are comparable; two pointers are equal if and only if they point to the same variable or both are `nil`.
```go
  var x, y int
  fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
```

It is perfectly safe for a function to return the address of a local variable. For instance, in the code below, the local variable v created by this particular call to f will remain in existence even after the call has returned, and the pointer p will still refer to it:
```go
  var p = f()
  func f() *int {
    v := 1
    return &v
  }
```

Each call of `f` returns a distinct value:
```go
  fmt.Println(f() == f()) // "false"
```

Because a pointer contains the address of a variable, passing a pointer argument to a function makes it possible for the function to update the variable that was indirectly passed. For example, this function increments the variable that its argument points to and returns the new value of the variable so it may be used in an expression:
```go
  func incr(p *int) int {
    *p++ // increments what p points to; does not change p
    return *p
  }
  v := 1
  incr(&v)              // side effect: v is now 2
  fmt.Println(incr(&v)) // "3" (and v is 3)
```

Each time we take the address of a variable or copy a pointer, we create new aliases or ways to identify the same variable. For example, `*p` is an alias for `v`. Pointer aliasing is useful because it allows us to access a variable without using its name, but this is a double-edged sword: to find all the statements that access a variable, we have to know all its aliases. It’s not just pointers that create aliases; aliasing also occurs when we copy values of other reference types like slices, maps, and channels, and even structs, arrays, and interfaces that contain these types.   

Pointers are key to the `flag` package, which uses a program’s command-line arguments to set the values of certain variables distributed throughout the program. To illustrate, this variation on the earlier `echo` command takes two optional flags: `-n` causes `echo` to omit the trailing newline that would normally be printed, and `-s sep` causes it to separate the output arguments by the contents of the string `sep` instead of the default single space. Since this is our fourth version, the package is called `gopl.io/ch2/echo4`. 
```go
// gopl.io/ch2/echo4
// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
```

The function `flag.Bool` creates a new flag variable of type `bool`. It takes three arguments: the name of the flag (`"n"`), the variable’s default value (`false`), and a message that will be printed if the user provides an invalid argument, an invalid flag, or `-h` or `-help`. Similarly, `flag.String` takes a name, a default value, and a message, and creates a string variable. The variables `sep` and `n` are pointers to the flag variables, which must be accessed indirectly as `*sep` and `*n`.

When the program is run, it must call `flag.Parse` before the flags are used, to update the flag variables from their default values. The non-flag arguments are available from `flag.Args()` as a slice of strings. If `flag.Parse` encounters an error, it prints a usage message and calls `os.Exit(2)` to terminate the program.

Let’s run some test cases on echo:
```bash
  $ go build gopl.io/ch2/echo4
  $ ./echo4 a bc def
  a bc def
  $ ./echo4 -s / a bc def
  a/bc/def
  $ ./echo4 -n a bc def
  a bc def$
  $ ./echo4 -help
  Usage of ./echo4:
    -n    
          omit trailing newline
    -s string
          separator (default " ")
```
### 2.3.3 The `new` Function

Another way to create a variable is to use the built-in function `new`. The expression `new(T)` creates an *unnamed variable* of type `T`, initializes it to the zero value of `T`, and returns its address, which is a value of type `*T`.
```go
  p := new(int)   // p, of type *int, points to an unnamed int variable
  fmt.Println(*p) // "0"
  *p = 2          // sets the unnamed int to 2
  fmt.Println(*p) // "2"
```

A variable created with new is no different from an ordinary local variable whose address is taken, except that there’s no need to invent (and declare) a dummy name, and we can use `new(T)` in an expression. Thus `new` is only a syntactic convenience, not a fundamental notion.
The two `newInt` functions below have identical behaviors.
```go
  func newInt() *int {
    return new(int)
  }
  func newInt() *int {
    var dummy int
    return &dummy
  }
```

Each call to `new` returns a distinct variable with a unique address:
```go

  p := new(int)
  q := new(int)
  fmt.Println(p == q) // "false"
```

There is one exception to this rule: two variables whose type carries no information and is therefore of size zero, such as `struct{}` or `[0]int`, may, depending on the implementation, have the same address.  

The `new` function is relatively rarely used because the most common unnamed variables are of struct types, for which the struct literal syntax (§4.4.1) is more flexible.  

Since new is a predeclared function, not a keyword, it’s possible to redefine the name for something else within a function, for example:
```go
  func delta(old, new int) int { return new - old }
```
Of course, within delta, the built-in `new` function is unavailable.

### 2.3.4 Lifetime of Variables
### 2.3.5 Tuple Assignment
### 2.3.6 Assignability
## 2.4. Assignments
## 2.5. Type Declarations 
## 2.6. Packages and Files 
### 2.6.1 Imports
### 2.6.2 Package Initialization
## 2.7. Scope

