# Notes

- Getting Started
  - Hello World
  - Numeral Systems
  - UTF-8
- Language Baics
  - Packages
  - Variables
    - Shorthand
    - Zero Value
  - Scope
    - Package Scope
    - Block Scope
    - ~~Order Matters~~
    - Variable Shadowing
    - ~~Same Package~~
  - Blank Identifier
  - Constants
    - Iota
  - Memory 
    - Showing Addresses
    - Using Addresses
  - Pointers
    - Referencing
    - Dereferencing
    - Using Pointers
  - Remainder
- Booleans
- Runes
- Control Flow
  - For Loop
  - Switch Statements
    - Switch
    - Fallthrough
    - Multiple Eval
    - No Expression
    - On Type
  - If Else 
  - Exercises
- Functions
  - Params vs. Args
    - Variadic
  - Returns
    - Named
    - Multiple
  - Closure
  - Callbacks
  - Recursion
  - Defer
  - Pass by Value
  - Anonymous Self-Executing
- Data Structures
  - Array
  - Slice
  - Map
  - Struct
  - Interfaces
- Go Routines
- Error Handling
- Testing


## Constants

```go
const p = "death & taxes"

const (
	pi       = 3.14
	language = "Go"
)

func main() {
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println(pi)
	fmt.Println(language)
}
```
