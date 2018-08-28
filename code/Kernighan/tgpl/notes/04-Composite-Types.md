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

An array is a fixed-length sequence of zero or more elements of a particular type. Because of their fixed length, arrays are rarely used directly in Go. Slices, which can grow and shrink, are much more versatile, but to understand slices we must understand arrays first.

Individual array elements are accessed with the conventional subscript notation, where subscripts run from zero to one less than the array length. The built-in function `len` returns the number of elements in the array.
```go
  var a [3]int             // array of 3 integers
  fmt.Println(a[0])        // print the first element
  fmt.Println(a[len(a)-1]) // print the last element, a[2]
  // Print the indices and elements.
  for i, v := range a {
      fmt.Printf("%d %d\n", i, v)
  }
  // Print the elements only.
  for _, v := range a {
      fmt.Printf("%d\n", v)
  }
```

By default, the elements of a new array variable are initially set to the zero value for the element type, which is 0 for numbers. We can use an *array literal* to initialize an array with a list of values:
```go
  var q [3]int = [3]int{1, 2, 3}
  var r [3]int = [3]int{1, 2}
  fmt.Println(r[2]) // "0"
```

In an array literal, if an ellipsis "`...`" appears in place of the length, the array length is determined by the number of initializers. The definition of q can be simplified to
```go
  q := [...]int{1, 2, 3}
  fmt.Printf("%T\n", q) // "[3]int"
```

The size of an array is part of its type, so `[3]int` and `[4]int` are different types. The size must be a constant expression, that is, an expression whose value can be computed as the program is being compiled.
```go
  q := [3]int{1, 2, 3}
  q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
```

As we’ll see, the literal syntax is similar for arrays, slices, maps, and structs. The specific form above is a list of values in order, but it is also possible to specify a list of index and value pairs, like this:
```go
  type Currency int
  const (
      USD Currency = iota
      EUR
      GBP
      RMB
  )
  symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"} 
  fmt.Println(RMB, symbol[RMB]) // "3 ¥"
```

In this form, indices can appear in any order and some may be omitted; as before, unspecified values take on the zero value for the element type. For instance,
```go
  r := [...]int{99: -1}
```
defines an array r with 100 elements, all zero except for the last, which has value −1.

If an array’s element type is *comparable* then the array type is comparable too, so we may directly compare two arrays of that type using the `==` operator, which reports whether all corresponding elements are equal. The `!=` operator is its negation.
```go
  a := [2]int{1, 2}
  b := [...]int{1, 2}
  c := [2]int{1, 3}
  fmt.Println(a == b, a == c, b == c) // "true false false"
  d := [3]int{1, 2}
  fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
```

As a more plausible example, the function `Sum256` in the `crypto/sha256` package produces the SHA256 cryptographic hash or *digest* of a message stored in an arbitrary byte slice. The digest has 256 bits, so its type is `[32]byte`. If two digests are the same, it is extremely likely that the two messages are the same; if the digests differ, the two messages are different. This program prints and compares the SHA256 digests of "x" and "X":
```go
// gopl.io/ch4/sha256
// The sha256 command computes the SHA256 hash (an array) of a string.
import (
  "crypto/sha256"
  "fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}
```

The two inputs differ by only a single bit, but approximately half the bits are different in the digests. Notice the `Printf` verbs: `%x` to print all the elements of an array or slice of bytes in hexadecimal, `%t` to show a boolean, and `%T` to display the type of a value.

When a function is called, a copy of each argument value is assigned to the corresponding parameter variable, so the function receives a copy, not the original. Passing large arrays in this way can be inefficient, and any changes that the function makes to array elements affect only the copy, not the original. In this regard, Go treats arrays like any other type, but this behavior is different from languages that implicitly pass arrays *by reference*.

Of course, we can explicitly pass a pointer to an array so that any modifications the function makes to array elements will be visible to the caller. This function zeroes the contents of a `[32]byte` array:
```go
  func zero(ptr *[32]byte) {
      for i := range ptr {
          ptr[i] = 0
      }
  }
```

The array literal `[32]byte{}` yields an array of 32 bytes. Each element of the array has the zero value for `byte`, which is zero. We can use that fact to write a different version of `zero`:
```go
  func zero(ptr *[32]byte) {
      *ptr = [32]byte{}
  }
```

Using a pointer to an array is efficient and allows the called function to mutate the caller’s variable, but arrays are still inherently inflexible because of their fixed size. The zero function will not accept a pointer to a `[16]byte` variable, for example, nor is there any way to add or remove array elements. For these reasons, other than special cases like SHA256’s fixed-size hash, arrays are seldom used as function parameters; instead, we use slices.

### Exercises
- **Exercise 4.1**: Write a function that counts the number of bits that are different in two SHA256 hashes. (See `PopCount` from Section2.6.2.)
- **Exercise 4.2**: Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.


## 4.2. Slices 

Slices represent variable-length sequences whose elements all have the same type. A slice type is written `[]T`, where the elements have type `T`; it looks like an array type without a size.

Arrays and slices are intimately connected. A slice is a lightweight data structure that gives access to a subsequence (or perhaps all) of the elements of an array, which is known as the slice’s *underlying array*. A slice has three components: a pointer, a length, and a capacity. The pointer points to the first element of the array that is reachable through the slice, which is not necessarily the array’s first element. The length is the number of slice elements; it can’t exceed the capacity, which is usually the number of elements between the start of the slice and the end of the underlying array. The built-in functions `len` and `cap` return those values.

Multiple slices can share the same underlying array and may refer to overlapping parts of that array. Figure 4.1 shows an array of strings for the months of the year, and two overlapping slices of it. The array is declared as
```go
  months := [...]string{1: "January", /* ... */, 12: "December"}
```
so January is `months[1]` and December is `months[12]`. Ordinarily, the array element at index 0 would contain the first value, but because months are always numbered from 1, we can leave it out of the declaration and it will be initialized to an empty string.

The *slice operator* `s[i:j]`, where `0 ≤ i ≤ j ≤ cap(s)`, creates a new slice that refers to elements `i` through `j-1` of the sequence `s`, which may be an array variable, a pointer to an array, or another slice. The resulting slice has `j-i` elements. If `i` is omitted, it’s `0`, and if `j` is omitted, it’s `len(s)`. Thus the slice `months[1:13]` refers to the whole range of valid months, as does the slice `months[1:]`; the slice `months[:]` refers to the whole array. Let’s define overlapping slices for the second quarter and the northern summer:

![Fig 4.1](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig4.1.png)

```go
  Q2 := months[4:7]
  summer := months[6:9]
  fmt.Println(Q2)     // ["April" "May" "June"]
  fmt.Println(summer) // ["June" "July" "August"]
```

June is included in each and is the sole output of this (inefficient) test for common elements:
```go
  for _, s := range summer {
      for _, q := range Q2 {
          if s == q {
              fmt.Printf("%s appears in both\n", s)
          }
      } 
  }
```
Slicing beyond `cap(s)` causes a panic, but slicing beyond `len(s)` extends the slice, so the result may be longer than the original:
```go
  fmt.Println(summer[:20]) // panic: out of range
  endlessSummer := summer[:5] // extend a slice (within capacity)
  fmt.Println(endlessSummer)  // "[June July August September October]"
```

As an aside, note the similarity of the substring operation on strings to the slice operator on `[]byte` slices. Both are written `x[m:n]`, and both return a subsequence of the original bytes, sharing the underlying representation so that both operations take constant time. The expression `x[m:n]` yields a string if x is a string, or a `[]byte` if `x` is a `[]byte`.

Since a slice contains a pointer to an element of an array, passing a slice to a function permits the function to modify the underlying array elements. In other words, copying a slice creates an *alias* (§2.3.2) for the underlying array. The function `reverse` reverses the elements of an `[]int` slice in place, and it may be applied to slices of any length.
```go
// gopl.io/ch4/rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
```

Here we reverse the whole array `a`:
```go
  a := [...]int{0, 1, 2, 3, 4, 5}
  reverse(a[:])
  fmt.Println(a) // "[5 4 3 2 1 0]"
```
A simple way to *rotate* a slice left by *n* elements is to apply the `reverse` function three times, first to the leading *n* elements, then to the remaining elements, and finally to the whole slice. (To rotate to the right, make the third call first.)
```go
  s := []int{0, 1, 2, 3, 4, 5}
  // Rotate s left by two positions.
  reverse(s[:2])
  reverse(s[2:])
  reverse(s)
  fmt.Println(s) // "[2 3 4 5 0 1]"
```

Notice how the expression that initializes the slice s differs from that for the array a. A *slice literal* looks like an array literal, a sequence of values separated by commas and surrounded by braces, but the size is not given. This implicitly creates an array variable of the right size and yields a slice that points to it. As with array literals, slice literals may specify the values in order, or give their indices explicitly, or use a mix of the two styles.

Unlike arrays, slices are not comparable, so we cannot use `==` to test whether two slices contain the same elements. The standard library provides the highly optimized `bytes.Equal` function for comparing two slices of bytes (`[]byte`), but for other types of slice, we must do the comparison ourselves:
```go
  func equal(x, y []string) bool {
      if len(x) != len(y) {
          return false
      }
      for i := range x {
          if x[i] != y[i] {
              return false
          } 
      }
      return true
  }
```

Given how natural this "deep" equality test is, and that it is no more costly at run time than the `==` operator for arrays of strings, it may be puzzling that slice comparisons do not also work this way. There are two reasons why deep equivalence is problematic. First, unlike array elements, the elements of a slice are indirect, making it possible for a slice to contain itself. Although there are ways to deal with such cases, none is simple, efficient, and most importantly, obvious.

Second, because slice elements are indirect, a fixed slice value may contain different elements at different times as the contents of the underlying array are modified. Because a hash table such as Go’s map type makes only shallow copies of its keys, it requires that equality for each key remain the same throughout the lifetime of the hash table. Deep equivalence would thus make slices unsuitable for use as map keys. For reference types like pointers and channels, the `==` operator tests *reference identity*, that is, whether the two entities refer to the same thing. An analogous "shallow" equality test for slices could be useful, and it would solve the problem with maps, but the inconsistent treatment of slices and arrays by the `==` operator would be confusing. The safest choice is to disallow slice comparisons altogether.

The only legal slice comparison is against `nil`, as in
```go
  if summer == nil { /* ... */ }
```

The zero value of a slice type is `nil`. A nil slice has no underlying array. The nil slice has length and capacity zero, but there are also non-nil slices of length and capacity zero, such as `[]int{}` or `make([]int, 3)[3:]`. As with any type that can have nil values, the nil value of a particular slice type can be written using a conversion expression such as `[]int(nil)`.
```go
  var s []int    // len(s) == 0, s == nil
  s = nil        // len(s) == 0, s == nil
  s = []int(nil) // len(s) == 0, s == nil
  s = []int{}    // len(s) == 0, s != nil
```

So, if you need to test whether a slice is empty, use `len(s) == 0`, not `s == nil`. Other than comparing equal to `nil`, a nil slice behaves like any other zero-length slice; `reverse(nil)` is perfectly safe, for example. Unless clearly documented to the contrary, Go functions should treat all zero-length slices the same way, whether nil or non-nil.   
The built-in function `make` creates a slice of a specified element type, length, and capacity. The capacity argument may be omitted, in which case the capacity equals the length.
```go
  make([]T, len)
  make([]T, len, cap) // same as make([]T, cap)[:len]
```

Under the hood, `make` creates an unnamed array variable and returns a slice of it; the array is accessible only through the returned slice. In the first form, the slice is a view of the entire array. In the second, the slice is a view of only the array’s first len elements, but its capacity includes the entire array. The additional elements are set aside for future growth.


### 4.2.1. The `append` Function

The built-in `append` function appends items to slices:
```go
  var runes []rune
  for _, r := range "Hello, 世界" {
      runes = append(runes, r)
  }
  fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
```
The loop uses append to build the slice of nine runes encoded by the string literal, although this specific problem is more conveniently solved by using the built-in conversion `[]rune("Hello, 世界")`.

The `append` function is crucial to understanding how slices work, so let’s take a look at what is going on. Here’s a version called `appendInt` that is specialized for `[]int` slices:
```go
// gopl.io/ch4/append
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
```

Each call to `appendInt` must check whether the slice has sufficient capacity to hold the new elements in the existing array. If so, it extends the slice by defining a larger slice (still within the original array), copies the element y into the new space, and returns the slice. The input `x` and the result `z` share the same underlying array.

If there is insufficient space for growth, `appendInt` must allocate a new array big enough to hold the result, copy the values from `x` into it, then append the new element `y`. The result `z` now refers to a different underlying array than the array that `x` refers to.

It would be straightforward to copy the elements with explicit loops, but it’s easier to use the built-in function `copy`, which copies elements from one slice to another of the same type. Its first argument is the destination and its second is the source, resembling the order of operands in an assignment like `dst = src`. The slices may refer to the same underlying array; they may even overlap. Although we don’t use it here, `copy` returns the number of elements actually copied, which is the smaller of the two slice lengths, so there is no danger of running off the end or overwriting something out of range.

For efficiency, the new array is usually somewhat larger than the minimum needed to hold `x` and `y`. Expanding the array by doubling its size at each expansion avoids an excessive number of allocations and ensures that appending a single element takes constant time on average. This program demonstrates the effect:
```go
  func main() {
      var x, y []int
      for i := 0; i < 10; i++ {
          y = appendInt(x, i)
          fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y) 
          x=y
      } 
  }
```

Each change in capacity indicates an allocation and a copy:
```
  0  cap=1   [0]
  1  cap=2   [0 1]
  2  cap=4   [0 1 2]
  3  cap=4   [0 1 2 3]
  4  cap=8   [0 1 2 3 4]
  5  cap=8   [0 1 2 3 4 5]
  6  cap=8   [0 1 2 3 4 5 6]
  7  cap=8   [0 1 2 3 4 5 6 7]
  8  cap=16  [0 1 2 3 4 5 6 7 8]
  9  cap=16  [0 1 2 3 4 5 6 7 8 9]
```
Let’stakeacloserlookatthei=3iteration. The slice `x` contains the three elements `[ 0 1 2 ]` but has capacity 4, so there is a single element of slack at the end, and `appendInt` of the element `3` may proceed without reallocating. The resulting slice y has length and capacity 4, and has the same underlying array as the original slice `x`, as Figure 4.2 shows.

![Fig 4.2](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig4.2.png)

On the next iteration, `i=4`, there is no slack at all, so `appendInt` allocates a new array of size 8, copies the four elements `[0 1 2 3]` of `x`, and appends 4, the value of `i`. The resulting slice y has a length of 5 but a capacity of 8; the slack of 3 will save the next three iterations from the need to reallocate. The slices `y` and `x` are views of different arrays. This operation is depicted in Figure 4.3.

![Fig 4.3](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig4.3.png)

The built-in append function may use a more sophisticated growth strategy than `appendInt`'s simplistic one. Usually we don’t know whether a given call to append will cause a reallocation, so we can’t assume that the original slice refers to the same array as the resulting slice, nor that it refers to a different one. Similarly, we must not assume that operations on elements of the old slice will (or will not) be reflected in the new slice. As a result, it’s usual to assign the result of a call to append to the same slice variable whose value we passed to append:
```go
  runes = append(runes, r)
```

Updating the slice variable is required not just when calling `append`, but for any function that may change the length or capacity of a slice or make it refer to a different underlying array. To use slices correctly, it’s important to bear in mind that although the elements of the underlying array are indirect, the slice’s pointer, length, and capacity are not. To update them requires an assignment like the one above. In this respect, slices are not "pure" reference types but resemble an aggregate type such as this struct:
```go
  type IntSlice struct {
      ptr      *int
      len, cap int
  }
```

Our `appendInt` function adds a single element to a slice, but the built-in `append` lets us add more than one new element, or even a whole slice of them.
```go
  var x []int
  x = append(x, 1)
  x = append(x, 2, 3)
  x = append(x, 4, 5, 6)
  x = append(x, x...) // append the slice x
  fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
```

With the small modification shown below, we can match the behavior of the built-in `append`. The ellipsis "..." in the declaration of `appendInt` makes the function *variadic*: it accepts any number of final arguments. The corresponding ellipsis in the call above to `append` shows how to supply a list of arguments from a slice. We’ll explain this mechanism in detail in Section 5.7.
```go
  func appendInt(x []int, y ...int) []int {
      var z []int
      zlen := len(x) + len(y)
      // ...expand z to at least zlen...
      copy(z[len(x):], y)
      return z
  }
```

The logic to expand `z`'s underlying array remains unchanged and is not shown.


### 4.2.2. In-Place Slice Techniques

Let’s see more examples of functions that, like `rotate` and `reverse`, modify the elements of a slice in place. Given a list of strings, the nonempty function returns the non-empty ones:
```go
// gopl.io/ch4/nonempty

```

The subtle part is that the input slice and the output slice share the same underlying array. This avoids the need to allocate another array, though of course the contents of `data` are partly overwritten, as evidenced by the second print statement:
```go
  data := []string{"one", "", "three"}
  fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
  fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
```

Thus we would usually write: `data = nonempty(data)`.

The nonempty function can also be written using `append`:
```go
  func nonempty2(strings []string) []string {
      out := strings[:0] // zero-length slice of original
      for _, s := range strings {
          if s != "" {
              out = append(out, s)
          }
      }
      return out
  }
```

Whichever variant we use, reusing an array in this way requires that at most one output value is produced for each input value, which is true of many algorithms that filter out elements of a sequence or combine adjacent ones. Such intricate slice usage is the exception, not the rule, but it can be clear, efficient, and useful on occasion.

A slice can be used to implement a stack. Given an initially empty slice stack, we can push a new value onto the end of the slice with append:
```go
  stack = append(stack, v) // push v
```
The top of the stack is the last element:
```go
  top := stack[len(stack)-1] // top of stack
```
and shrinking the stack by popping that element is
```go
  stack = stack[:len(stack)-1] // pop
```

To remove an element from the middle of a slice, preserving the order of the remaining elements, use copy to slide the higher-numbered elements down by one to fill the gap:
```go
  func remove(slice []int, i int) []int {
      copy(slice[i:], slice[i+1:])
      return slice[:len(slice)-1]
  }
      
  func main() {
      s := []int{5, 6, 7, 8, 9}
      fmt.Println(remove(s, 2)) // "[5 6 8 9]"
  }
```
And if we don’t need to preserve the order, we can just move the last element into the gap:
```go
  func remove(slice []int, i int) []int {
      slice[i] = slice[len(slice)-1]
      return slice[:len(slice)-1]
  }

  func main() {
      s := []int{5, 6, 7, 8, 9}
      fmt.Println(remove(s, 2)) // "[5 6 9 8]
  }
```

#### Exercises
- **Exercise 4.3**: Rewrite `reverse` to use an array pointer instead of a slice.
- **Exercise 4.4**: Write a version of rotate that operates in a single pass.
- **Exercise 4.5**: Write an in-place function to eliminate adjacent duplicates in a `[]string` slice.
- **Exercise 4.6**: Write an in-place function that squashes each run of adjacent Unicode spaces (see `unicode.IsSpace`) in a UTF-8-encoded `[]byte` slice into a single ASCII space.
- **Exercise 4.7**: Modify reverse to reverse the characters of a `[]byte` slice that represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?


## 4.3. Maps 











## 4.4. Structs 
### 4.4.1. Struct Literals
### 4.4.2. Comparing Structs
### 4.4.3. Struct Embedding and Anonymous Fields
## 4.5. JSON 
## 4.6. Text and HTML Templates

