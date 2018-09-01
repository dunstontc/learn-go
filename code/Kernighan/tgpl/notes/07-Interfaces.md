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

Throughout the book, we’ve been using two similar functions for string formatting: fmt.Printf, which writes the result to the standard output (a file), and fmt.Sprintf, which returns the result as a string. It would be unfortunate if the hard part, formatting the result, had to be duplicated because of these superficial differences in how the result is used. Thanks to interfaces, it does not. Both of these functions are, in effect, wrappers around a third function, fmt.Fprintf, that is agnostic about what happens to the result it computes:


## 7.2. Interface Types 
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
