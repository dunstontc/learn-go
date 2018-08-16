package main

import "fmt"

func main() {
	rem := 7.24
	fmt.Printf("%T\n", rem)      // float64
	fmt.Printf("%T\n", int(rem)) // int
}
