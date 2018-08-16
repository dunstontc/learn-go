package main

import "fmt"

func main() {
	rem := 7.24
	fmt.Printf("%T\n", rem)      // float64
	fmt.Printf("%T\n", int(rem)) // int

	var val interface{} = 7
	fmt.Printf("%T\n", val) // int
	// fmt.Printf("%T\n", int(val)) // cannot convert val (type interface {}) to type int: need type assertion
	fmt.Printf("%T\n", val.(int)) // int
}
