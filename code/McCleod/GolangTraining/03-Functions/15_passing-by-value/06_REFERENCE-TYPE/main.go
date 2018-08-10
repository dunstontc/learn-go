package main

import "fmt"

func main() {
	m := make([]string, 1)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Clay]
}

func changeMe(z []string) {
	z[0] = "Clay"
	fmt.Println(cap(z)) // [Clay]
}
