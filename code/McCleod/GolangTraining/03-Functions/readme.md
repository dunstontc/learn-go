# Functions

- parameters vs. arguments
  - params
    multiple “variadic” params
  - args
    - multiple “variadic” args
- returns
  - multiple returns
  - named returns - yuck!
- review
  - func expressions
  - closure 
- callbacks
- recursion
- defer
- anonymous self-executing functions
- pass by value
  - reference types


## Parameters v. Arguments

## Returns

## Callbacks

```go
package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}
```

## Recursion

```go
package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}
```

## Defer
