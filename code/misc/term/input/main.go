// https://gist.github.com/eduncan911/37a351731b0ddeeeb2d0
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print(`
Which example do you want to run?
  1) fmt.Scan(...)
  2) bufio.Reader.ReadString(...)
  3) bufio.Reader.ReadByte(...)
  4) bufio.Reader.ReadRune()
Please enter 1..5 and press ENTER:
`)

	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		return
	}

	switch result {

	case '1':
		runScan()
		break

	case '2':
		runReadString()
		break

	case '3':
		runReadByte()
		break

	case '4':
		runReadRune()
		break

	default:
		return
	}

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

// if using ReadString() a lot, consider using constants
const inputdelimiter = '\n'

func runReadString() {

	fmt.Print("\nEnter your Full Name: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(inputdelimiter)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)

	fmt.Println(input)
	fmt.Println("Exiting program.")

}

func runReadByte() {

	fmt.Print("\nContinue? [Y/N] ")

	reader := bufio.NewReader(os.Stdin)
	c, err := reader.ReadByte()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("You entered: %q\n", c)

	if c == []byte("Y")[0] || c == []byte("y")[0] {
		fmt.Println("Thank you for pressing Y to continue!")
	} else {
		fmt.Println("No? Ok, we'll exit.")
	}
}

func runReadRune() {

	fmt.Print("\nContinue? [Y/N] ")

	reader := bufio.NewReader(os.Stdin)
	c, num, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("You entered: %q\n", c)
	fmt.Println("The size entered: ", num)

	if c == 'y' || c == 'Y' {
		fmt.Println("Thank you for pressing Y to continue!")
	} else {
		fmt.Println("No? Ok, we'll exit.")
	}
}
