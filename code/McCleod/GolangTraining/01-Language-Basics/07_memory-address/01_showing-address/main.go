package main

import "fmt"

func main() {

	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("typeof(a) = %T \n", a)
	fmt.Printf("typeof(&a) = %T \n", &a)
	fmt.Printf("%d \n", &a)
}
