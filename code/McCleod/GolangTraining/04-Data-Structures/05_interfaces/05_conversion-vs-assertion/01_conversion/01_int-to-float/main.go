package main

import "fmt"

func main() {
	var x = 12
	var y = 12.1230123
	fmt.Println(y + float64(x)) // 24.1230123
	// conversion: int to float64
}
