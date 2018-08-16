package main

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	blue := color.New(color.FgBlue).SprintFunc()

	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", blue(i))
	}
	wg.Done()
}

func bar() {
	yellow := color.New(color.FgYellow).SprintFunc()

	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", yellow(i))
	}
	wg.Done()
}
