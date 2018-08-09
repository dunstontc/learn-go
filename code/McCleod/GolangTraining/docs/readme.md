# Golang

## Programming

### Numeral Systems

|         |                |                        |               |
| ------- | -------------- | ---------------------- | ------------- |
| 1 bit   | 1 binary digit | 1 bit                  | 1             |
| 8 bits  | 1 byte         | 8 bits                 | 8             |
| 1000 b  | 1 KiloByte     | 8,000 bits             | 8 Thousand    |
| 1000 kb | 1 MegaBtye     | 8,000,000 bits         | *8 Million*   |
| 1000 mb | 1 GigaByte     | 8,000,000,000 bits     | 8 Billion     |
| 1000 gb | 1 TeraByte     | 8,000,000,000,000 bits | *8 Trillion*  |
| 1000 tb | 1 PetaByte     |                        | 8 Quadrillion |
| 1000 pb | 1 Exabyte      |                        | 8 Quintillion |
| 1000 eb | 1 Zettabyte    |                        | 8 Sextillion  |
| 1000 zb | 1 Yottabyte    |                        | 8 Septillion  |


|   Decimal    |          |                   |               |           |          |        |        |
| ------------ | -------- | ----------------- | ------------- | --------- | -------- | ------ | ------ |
| ten millions | millions | hundred thousands | ten thousands | thousands | hundreds | tens   | ones   |
| $10^7$       | $10^6$   | $10^5$            | $10^4$        | $10^3$    | $10^2$   | $10^1$ | $10^0$ |
|              |          |                   |               |           |          | 4      | 2      |

| Binary |       |       |       |       |       |       |       |
| ------ | ----- | ----- | ----- | ----- | ----- | ----- | ----- |
| 128's  | 64's  | 32's  | 16's  | 8's   | 4's   | 2's   | ones  |
| $2^7$  | $2^6$ | $2^5$ | $2^4$ | $2^3$ | $2^2$ | $2^1$ | $2^0$ |
|        |       | 1     |  0    | 1     | 0     | 1     | 0     |

| Hexadecimal |        |        |          |         |        |        |        |
| ----------- | ------ | ------ | -------- | ------- | ------ | ------ | ------ |
|             |        |        | 65,536's | 3,096's | 256's  | 16's   | ones   |
| $16^7$      | $16^6$ | $16^5$ | $16^4$   | $16^3$  | $16^2$ | $16^1$ | $16^0$ |
|             |        |        |          |         |        |        |        |


0,1,2,3,4,5,6,7,8,9,a,b,c,d,e,f



## Misc

|   Type    |                                                                                                                                               Description                                                                                                                                                |    Name     | Zero Value |
| --------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- | ---------- |
| Method    |                                                                                                                                                                                                                                                                                                          | `*T` or `T` |            |
| Boolean   | A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false.                                                                                                                                                                                           | `bool`      | `false`    |
| Numeric   | A numeric type represents sets of integer or floating-point values.                                                                                                                                                                                                                                      |             | `0`        |
| String    | A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. <br> Strings are immutable: once created, it is impossible to change the contents of a string.                                                                                                | `string`    | `""`       |
| Array     | An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length and is never negative.                                                                                                                                                |             |            |
| Slice     | A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The value of an uninitialized slice is nil.                                    |             |            |
| Struct    | A struct is a sequence of named elements, called fields, each of which has a name and a type. <br> Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique.                                                        | `struct {}` |            |
| Pointer   | A pointer type denotes the set of all pointers to variables of a given type, called the base type of the pointer. <br> The value of an uninitialized pointer is nil.                                                                                                                                     | `*Point`    |            |
| Function  | A function type denotes the set of all functions with the same parameter and result types.                                                                                                                                                                                                               |             |            |
| Interface | An interface type specifies a method set called its interface. A variable of interface type can store a value of any type with a method set that is any superset of the interface. Such a type is said to implement the interface. <br> The value of an uninitialized variable of interface type is nil. |             |            |
| Map       | A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type.                                                                                                                                                      |             |            |
| Channel   | A channel provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type. <br> The value of an uninitialized channel is nil.                                                                                                       |             |            |
| Nil       |                                                                                                                                                                                                                                                                                                          | `nil`       |            |


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


## Playgrounds

- [The Go Playground](https://play.golang.org/)
- [The Go Play Space](https://goplay.space/)
  - [decimal](https://goplay.space/#VaqXxWCQBiw)
  - [binary](https://goplay.space/#IuZDvLL4EUu)
  - [hex](https://goplay.space/#4vWR8_1Df3S)
  - [loop](https://goplay.space/#RI8kukvUrgb)
