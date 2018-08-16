package main

import "fmt"

func main() {
	var val interface{} = 7
	fmt.Println(val + 6)
	// ./main.go:7:18: invalid operation: val + 6 (mismatched types interface {} and int)
}
