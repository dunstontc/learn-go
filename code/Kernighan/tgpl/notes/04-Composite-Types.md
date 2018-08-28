# Chapter 4: Composite Types

<!-- TOC -->

- [4.1. Arrays](#41-arrays)
- [4.2. Slices](#42-slices)
  - [4.2.1. The `append` Function](#421-the-append-function)
  - [4.2.2. In-Place Slice Techniques](#422-in-place-slice-techniques)
- [4.3. Maps](#43-maps)
- [4.4. Structs](#44-structs)
  - [4.4.1. Struct Literals](#441-struct-literals)
  - [4.4.2. Comparing Structs](#442-comparing-structs)
  - [4.4.3. Struct Embedding and Anonymous Fields](#443-struct-embedding-and-anonymous-fields)
- [4.5. JSON](#45-json)
- [4.6. Text and HTML Templates](#46-text-and-html-templatesp)

<!-- /TOC -->

In Chapter 3 we discussed the basic types that serve as building blocks for data structures in a Go program; they are the atoms of our universe. In this chapter, we’ll take a look at *composite types*, the molecules created by combining the basic types in various ways. We’ll talk about four such types—arrays, slices, maps, and structs—and at the end of the chapter, we’ll show how structured data using these types can be encoded as and parsed from JSON data and used to generate HTML from templates.

Arrays and structs are *aggregate types*; their values are concatenations of other values in memory. Arrays are homogeneous—their elements all have the same type—whereas structs are heterogeneous. Both arrays and structs are fixed size. In contrast, slices and maps are dynamic data structures that grow as values are added.

## 4.1. Arrays 
## 4.2. Slices 
![Fig 4.1](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig4.1.png)
### 4.2.1. The `append` Function
![Fig 4.2](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig4.2.png)
![Fig 4.3](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig4.3.png)
### 4.2.2. In-Place Slice Techniques
## 4.3. Maps 
## 4.4. Structs 
### 4.4.1. Struct Literals
### 4.4.2. Comparing Structs
### 4.4.3. Struct Embedding and Anonymous Fields
## 4.5. JSON 
## 4.6. Text and HTML Templates

