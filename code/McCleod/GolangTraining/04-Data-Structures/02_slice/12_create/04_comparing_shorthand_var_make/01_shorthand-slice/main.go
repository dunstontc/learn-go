package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}
	student[0] = "Todd" // panic: runtime error: index out of range
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
