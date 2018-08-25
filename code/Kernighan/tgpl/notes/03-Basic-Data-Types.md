# Chapter 3: Basic Data Types

<!-- TOC -->

- [3.1. Integers](#31-integers)
- [3.2. Floating-Point Numbers](#32-floating-point-numbers)
- [3.3. Complex Numbers](#33-complex-numbers)
- [3.4. Booleans](#34-booleans)
- [3.5. Strings](#35-strings)
  - [3.5.1 String Literals](#351-string-literals)
  - [3.5.2 Unicode](#352-unicode)
  - [3.5.3 UTF-8](#353-utf-8)
  - [3.5.4 Strings and Byte Slices](#354-strings-and-byte-slices)
  - [3.5.5 Conversions between Strings and Numbers](#355-conversions-between-strings-and-numbers)
- [3.6. Constants](#36-constants)
  - [3.6.1 The Constant Generator `iota`](#361-the-constant-generator-iota)
  - [3.6.2 Untyped Constants](#362-untyped-constants)

<!-- /TOC -->


It’s all bits at the bottom, of course, but computers operate fundamentally on fixed-size numbers called *words*, which are interpreted as integers, floating-point numbers, bit sets, or memory addresses, then combined into larger aggregates that represent packets, pixels, portfolios, poetry, and everything else. Go offers a variety of ways to organize data, with a spectrum of data types that at one end match the features of the hardware and at the other end provide what programmers need to conveniently represent complicated data structures.

Go’s types fall into four categories: *basic types*, *aggregate types*, *reference types*, and *interface types*. Basic types, the topic of this chapter, include numbers, strings, and booleans. Aggregate types—arrays (§4.1) and structs (§4.4)—form more complicated data types by combining val- ues of several simpler ones. Reference types are a diverse group that includes pointers (§2.3.2), slices (§4.2), maps (§4.3), functions (Chapter 5), and channels (Chapter 8), but what they have in common is that they all refer to program variables or state indirectly, so that the effect of an operation applied to one reference is observed by all copies of that reference. Finally, we’ll talk about interface types in Chapter 7.

## 3.1. Integers 

Go’s numeric data types include several sizes of integers, floating-point numbers, and complex numbers. Each numeric type determines the size and signedness of its values. Let’s begin with integers.

Go provides both signed and unsigned integer arithmetic. There are four distinct sizes of signed integers—8, 16, 32, and 64 bits—represented by the types `int8`, `int16`, `int32`, and `int64`, and corresponding unsigned versions `uint8`, `uint16`, `uint32`, and `uint64`.

There are also two types called just `int` and `uint` that are the natural or most efficient size for signed and unsigned integers on a particular platform; int is by far the most widely used numeric type. Both these types have the same size, either 32 or 64 bits, but one must not make assumptions about which; different compilers may make different choices even on identical hardware.

The type `rune` is an synonym for `int32` and conventionally indicates that a value is a Unicode code point. The two names may be used interchangeably. Similarly, the type `byte` is an synonym for `uint8`, and emphasizes that the value is a piece of raw data rather than a small numeric quantity.

## 3.2. Floating-Point Numbers 
## 3.3. Complex Numbers 
## 3.4. Booleans 
## 3.5. Strings 
### 3.5.1 String Literals
### 3.5.2 Unicode
### 3.5.3 UTF-8
### 3.5.4 Strings and Byte Slices
### 3.5.5 Conversions between Strings and Numbers
## 3.6. Constants
### 3.6.1 The Constant Generator `iota`
### 3.6.2 Untyped Constants

