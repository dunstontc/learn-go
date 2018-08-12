<!--
video #80
Slice
definition
A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array. 
The value of an uninitialized slice is nil. 
it is a reference type
Like arrays, slices are indexable and have a length. 
The length of a slice s can be discovered by the built-in function len; 
Unlike arrays, slices are dynamic
their length may change during execution. 
The elements can be addressed by integer indices 0 through len(s)-1.
A slice, once initialized, is always associated with an underlying array that holds its elements. 
it is a reference type
The array underlying a slice may extend past the end of the slice. 
Capacity is a measure of that extent: 
it is the sum of the length of the slice and the length of the array beyond the slice; 
The capacity of a slice a can be discovered using the built-in function cap(a). 
make
A new, initialized slice value for a given element type T is made using the built-in function make, which takes a slice type and parameters specifying the length and optionally the capacity. 
A slice created with make always allocates a new, hidden array to which the returned slice value refers. 
make([]T, length, capacity) 
make([]int, 50, 100) 
same as this: new([100]int)[0:50] 
Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. (multi-dimensional slices)
a basic slice 

video #81
Slice Examples
length and capacity
a great example
index out of range errors
appending items to slices
deleting items from slices


video #82
More Slice Examples
multidimensional slice
incrementing a slice


video #83
Creating A Slice
shorthand
var
sets slice to zero value which is nil
make


video #84
Incrementing A Slice Item
incrementing a slice item
review of slices
len, cap, underlying array, append

video #84_02
Section Review
definition
a list of values of a certain Type
internals
reference type
pointer, len, cap
built on-top of an array
another way to say it: “points to an array”
The value of an uninitialized slice is nil. 
because it is a reference type
A slice, once initialized, is always associated with an underlying array that holds its elements. 
slices are dynamic (unlike arrays)
their length may change during execution. 
The array underlying a slice may extend past the end of the slice. 
Capacity is a measure of that extent: 
The capacity of a slice a can be discovered using the built-in function cap(a). 
make
A slice created with make always allocates a new, hidden array to which the returned slice value refers. 
make([]T, length, capacity) 
make([]int, 50, 100) 
same as this: new([100]int)[0:50] 
Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. (multi-dimensional slices)
index out of range errors
appending items to slices
access by index if the index is less than the length of the slice less one
 0 through len(s)-1.
deleting items from slices
mySlice = append(mySlice[:2], mySlice[3:]...)
incrementing a slice
mySlice[0]++
creating a slice
shorthand
student := []string{}
var
sets slice to zero value which is nil
var student []string
make
student := make([]string, 35)
-->

# Slices

## Slicing a Slice

```go
func main() {

	var results []int
	fmt.Println(results)

	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])  // slicing a slice
	fmt.Println(mySlice[2])    // index access; accessing by index
	fmt.Println("myString"[2]) // index access; accessing by index
}
```

## Make

```go
func main() {

	customerNumber := make([]int, 3)
	// 3 is length & capacity
	// length - number of elements referred to by the slice
	// capacity - number of elements in the underlying array
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"

	fmt.Println(greeting[2])
}
```

## Append

```go
func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)
}
```

## Delete

```go
func main() {

	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)

}
```

## Create

### *Shorthand*
```go
func main() {
	student := []string{}
	students := [][]string{}
	fmt.Println(student)        // []
	fmt.Println(students)       // []
	fmt.Println(student == nil) // false
}
// requires the use of append
```

### *Var*
```go
func main() {
	var student []string
	var students [][]string
	fmt.Println(student)        // []
	fmt.Println(students)       // []
	fmt.Println(student == nil) // true
}
// requires the use of append
```

### *Make*
```go
func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	fmt.Println(student)        // [
	fmt.Println(students)       // [[] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] []]
	fmt.Println(student == nil) // false
}
// Length & Capaciity get set with `make()`
```


## Multidimensional
