package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 12
	y := "I have this many: " + strconv.Itoa(x)
	fmt.Println(y) // I have this many: 12
	// fmt.Println("I have this many: ", strconv.Itoa(x), x)
}
