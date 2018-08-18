/*Package main provides lines. */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Print(hide())
	runReadString()
	// fmt.Print(show())

}

func runScan() {

	// you must declare your var, and pass the pointer into Scan() below
	var input string

	fmt.Print("\nEnter some text and press enter: ")

	// using fmt.Scan, we can read single words in ascii string
	num, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(input)
	fmt.Println(num)
}

func runReadString() {

	fmt.Print("Enter your Full Name: ")
	fmt.Print("\r")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)

	fmt.Printf("Nice to meet you %s!", input)
	// fmt.Println("")
	// fmt.Println("Exiting program.")

}

/* https://github.com/ahmetb/go-cursor/blob/master/cursor.go */

// Show returns ANSI escape sequence to show the cursor
func show() string {
	return fmt.Sprintf("%c%s", '\x1b', fmt.Sprintf("[?25h"))
}

// Hide returns ANSI escape sequence to hide the cursor
func hide() string {
	return fmt.Sprintf("%c%s", '\x1b', fmt.Sprintf("[?25l"))
}
