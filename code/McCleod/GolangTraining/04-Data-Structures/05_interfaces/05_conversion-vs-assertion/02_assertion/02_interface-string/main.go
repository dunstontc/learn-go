package main

import "fmt"

func main() {
	var name interface{} = "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str) // string
	} else {
		fmt.Printf("value is not a string\n")
	}
}
