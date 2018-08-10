package main

import "fmt"

func main() {
	var x [58]int

	fmt.Println(x)      // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(x)) // 58
	fmt.Println(x[42])  // 0
	x[42] = 777
	fmt.Println(x[42]) // 777
}
