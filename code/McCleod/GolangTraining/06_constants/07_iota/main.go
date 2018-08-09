package main

import "fmt"

const (
	_  = iota             // 0
	kb = 1 << (iota * 10) // 1 << (1 * 10)
	mb = 1 << (iota * 10) // 1 << (2 * 10)
	gb = 1 << (iota * 10) // 1 << (3 * 10)
	tb = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("KB: %b\t", kb)
	fmt.Printf("KB: %d\n", kb)
	fmt.Printf("MB: %b\t", mb)
	fmt.Printf("MB: %d\n", mb)
	fmt.Printf("GB: %b\t", gb)
	fmt.Printf("GB: %d\n", gb)
	fmt.Printf("TB: %b\t", tb)
	fmt.Printf("TB: %d\n", tb)
}
