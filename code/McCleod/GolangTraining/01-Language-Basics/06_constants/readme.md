# Constants

- *constant*
  - a simple, unchanging value
- iota's
  - creating constants values for:
    - KB
    - MB
    - GB
    - TB

## More on Constants

- a parallel type system
  - C / C++ has problems with a lack of strict typing
  - in Go, there is no mixing of numeric types
- there are TYPED and UNTYPED constants
  - const hello = "Hello, World"
  - const typedHello string = "Hello, World"
- UNTYPED constant
  -  a constant value that does not yet have a fixed type
    - a “kind”
    - not yet forced to obey the strict rules that prevent combining differently typed values
- It is this notion of an untyped constant that makes it possible for us to use constants in Go with great freedom.
- This is useful, for instance
  - what is the type of 42?
    - int?
    - uint?
    - float64?
  - if we didn't have UNTYPED constants (constants of a kind), then we would have to do conversion on every literal value we used
    - and that would suck
