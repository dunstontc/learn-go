package main

import "fmt"

func main() {
	var x rune = 'a' // rune is an alias for int32; normally omitted in this statement
	var y int32 = 'b'
	fmt.Println(x)         // 97
	fmt.Println(y)         // 98
	fmt.Println(string(x)) // a
	fmt.Println(string(y)) // b
	// conversion: rune to string
}
