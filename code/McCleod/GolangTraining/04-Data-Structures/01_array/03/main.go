package main

import "fmt"

func main() {
	var x [256]int

	// fmt.Println(len(x)) // 256
	// fmt.Println(x[42])  // 0
	for i := 0; i < 256; i++ {
		x[i] = i
	}
	fmt.Println("|    |     |        |")
	fmt.Println("|----|-----|--------|")
	for i, v := range x {
		fmt.Printf("| %02d | %T | %6b |\n", v, v, v)
		if i > 50 {
			break
		}
	}
}
