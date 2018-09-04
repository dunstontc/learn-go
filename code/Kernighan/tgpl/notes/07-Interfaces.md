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

Many object-oriented languages have some notion of interfaces, but what makes Go’s interfaces so distinctive is that they are *satisfied implicitly*. In other words, there’s no need to declare all the interfaces that a given concrete type satisfies; simply possessing the necessary methods is enough. This design lets you create new interfaces that are satisfied by existing concrete types without changing the existing types, which is particularly useful for types defined in packages that you don’t control.

In this chapter, we’ll start by looking at the basic mechanics of interface types and their values. Along the way, we’ll study several important interfaces from the standard library. Many Go programs make as much use of standard interfaces as they do of their own ones. Finally, we’ll look at *type assertions* (§7.10) and *type switches* (§7.13) and see how they enable a different kind of generality.


## 7.1. Interfaces as Contracts

All the types we’ve looked at so far have been *concrete types*. A concrete type specifies the exact representation of its values and exposes the intrinsic operations of that representation, such as arithmetic for numbers, or indexing, append, and range for slices. A concrete type may also provide additional behaviors through its methods. When you have a value of a concrete type, you know exactly what it is and what you can do with it.

There is another kind of type in Go called an *interface type*. An interface is an *abstract type*. It doesn’t expose the representation or internal structure of its values, or the set of basic operations they support; it reveals only some of their methods. When you have a value of an interface type, you know nothing about what it is; you know only what it can do, or more precisely, what behaviors are provided by its methods.

Throughout the book, we’ve been using two similar functions for string formatting: `fmt.Printf`, which writes the result to the standard output (a file), and `fmt.Sprintf`, which returns the result as a string. It would be unfortunate if the hard part, formatting the result, had to be duplicated because of these superficial differences in how the result is used. Thanks to interfaces, it does not. Both of these functions are, in effect, wrappers around a third function, `fmt.Fprintf`, that is agnostic about what happens to the result it computes:
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

The first parameter of `Fprintf` is not a file either. It’s an `io.Writer`, which is an interface type with the following declaration:
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

Let’s test this out using a new type. The Write method of the *ByteCounter type below merely counts the bytes written to it before discarding them. (The conversion is required to make the types of `len(p)` and `*c` match in the `+=` assignment statement.)
```go
// gopl.io/ch7/bytecounter

```
Since `*ByteCounter` satisfies the io.Writer contract, we can pass it to `Fprintf`, which does its string formatting oblivious to this change; the `ByteCounter` correctly accumulates the length of the result.
```go

```
Besides io.Writer, there is another interface of great importance to the fmt package. Fprintf and Fprintln provide a way for types to control how their values are printed. In Section 2.5, we defined a String method for the Celsius type so that temperatures would print as "100°C", and in Section 6.5 we equipped *IntSet with a String method so that sets would be rendered using traditional set notation like "{1 2 3}". Declaring a String method makes a type satisfy one of the most widely used interfaces of all, fmt.Stringer:
```go

```
We’ll explain how the fmt package discovers which values satisfy this interface in Section 7.10.

### Exercises
- **Exercise 7.1**: Using the ideas from `ByteCounter`, implement counters for words and for lines. You will find `bufio.ScanWords` useful.
- **Exercise 7.2**: Write a function `CountingWriter` with the signature below that, given an `io.Writer`, returns a new Writer that wraps the original, and a pointer to an int64 variable that at any moment contains the number of bytes written to the new `Writer`.
```go
  func CountingWriter(w io.Writer) (io.Writer, *int64)
```
- **Exercise 7.3**: Write a `String` method for the `*tree` type in `gopl.io/ch4/treesort` (§4.4) that reveals the sequence of values in the tree.


## 7.2. Interface Types 

An interface type specifies a set of methods that a concrete type must possess to be considered an instance of that interface.

The io.Writer type is one of the most widely used interfaces because it provides an abstraction of all the types to which bytes can be written, which includes files, memory buffers, network connections, HTTP clients, archivers, hashers, and so on. The `io` package defines many other useful interfaces. A `Reader` represents any type from which you can read bytes, and a `Closer` is any value that you can close, such as a file or a network connection. (By now you’ve probably noticed the naming convention for many of Go’s single-method interfaces.)
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
## 7.4. Parsing Flags with flag.Value 
## 7.5. Interface Values
## 7.6. Sorting with sort.Interface 
## 7.7. The http.Handler Interface 
## 7.8. The error Interface 
## 7.9. Example: Expression Evaluator 
## 7.10. Type Assertions 
## 7.11. Discriminating Errors with Type Assertions 
## 7.12. Querying Behaviors with Interface Type Assertions 
## 7.13. Type Switches 
## 7.14. Example: Token-Based XML Decoding 
## 7.15. A Few Words of Advice
