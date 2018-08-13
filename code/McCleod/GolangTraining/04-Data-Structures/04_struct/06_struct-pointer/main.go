package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"James", 20}

	fmt.Println(p1)        // &{James 20}
	fmt.Printf("%T\n", p1) // *main.person
	fmt.Println(p1.name)   // James
	fmt.Println(p1.age)    // 20
}
