# Golang

## Links
- [GoesToEleven/golang-web-dev](https://github.com/GoesToEleven/golang-web-dev)
- [GoesToEleven/GolangTraining](https://github.com/GoesToEleven/GolangTraining)
- [Course outline](https://docs.google.com/document/d/1nt5bYAAS5sTVF6tpLaFLDHQzo5BNkcr4b507fg3ZPwM/edit)
- [Slides](https://drive.google.com/drive/folders/0B22KXlqHz6ZNfjNXTzk1U3JHUkJ6VjJ3dnJKNzVtNjRUM3Q2WFNqWGI2Q3RadERqUlVrOEU)
- [Spec](https://golang.org/ref/spec)
- [Effective Go](https://golang.org/doc/effective_go.html)

## Misc

|   Type    |                                                                                                                              Description                                                                                                                              |    Name     |
| --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| Method    |                                                                                                                                                                                                                                                                       | `*T` or `T` |
| Boolean   |                                                                                                                                                                                                                                                                       | `bool`      |
| Numeric   | A numeric type represents sets of integer or floating-point values.                                                                                                                                                                                                   |             |
| String    | A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. Strings are immutable: once created, it is impossible to change the contents of a string.                                                                  | `string`    |
| Array     | An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length and is never negative.                                                                                                             |             |
| Slice     | A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The value of an uninitialized slice is nil. |             |
| Struct    | A struct is a sequence of named elements, called fields, each of which has a name and a type. <br> Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique.                          |             |
| Pointer   | A pointer type denotes the set of all pointers to variables of a given type, called the base type of the pointer. <br> The value of an uninitialized pointer is nil.                                                                                                       |             |
| Function  | A function type denotes the set of all functions with the same parameter and result types.                                                                                                                                                                            |             |
| Interface |                                                                                                                                                                                                                                                                       |             |
| Map       |                                                                                                                                                                                                                                                                       |             |
| Channel   |                                                                                                                                                                                                                                                                       |             |
| Nil       |                                                                                                                                                                                                                                                                       | `nil`       |


|     Type     |                                     Description                                     |                    Range                    |
| ------------ | ----------------------------------------------------------------------------------- | ------------------------------------------- |
| `uint8`      | the set of all unsigned  8-bit integers                                             | 0 to 255                                    |
| `uint16`     | the set of all unsigned 16-bit integers                                             | 0 to 65535                                  |
| `uint32`     | the set of all unsigned 32-bit integers                                             | 0 to 4294967295                             |
| `uint64`     | the set of all unsigned 64-bit integers                                             | 0 to 18446744073709551615                   |
|              |                                                                                     |                                             |
| `int8`       | the set of all signed  8-bit integers                                               | -128 to 127                                 |
| `int16`      | the set of all signed 16-bit integers                                               | -32768 to 32767                             |
| `int32`      | the set of all signed 32-bit integers                                               | -2147483648 to 2147483647                   |
| `int64`      | the set of all signed 64-bit integers                                               | -9223372036854775808 to 9223372036854775807 |
|              |                                                                                     |                                             |
| `float32`    | the set of all IEEE-754 32-bit floating-point numbers                               |                                             |
| `float64`    | the set of all IEEE-754 64-bit floating-point numbers                               |                                             |
|              |                                                                                     |                                             |
| `complex64`  | the set of all complex numbers with float32 real and imaginary parts                |                                             |
| `complex128` | the set of all complex numbers with float64 real and imaginary parts                |                                             |
|              |                                                                                     |                                             |
| `byte`       | alias for uint8                                                                     |                                             |
| `rune`       | alias for int32                                                                     |                                             |
|              |                                                                                     |                                             |
| `uint`       | either 32 or 64 bits                                                                |                                             |
| `int`        | same size as uint                                                                   |                                             |
| `uintptr`    | an unsigned integer large enough to store the uninterpreted bits of a pointer value |                                             |