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
- pass by value
  - reference types
- Immediately Invoked Function Expressions (iifes)


## Parameters v. Arguments

## Returns

## Callbacks

```go
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

```go
func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello()
}
```


## IIFEs

```go
func main() {
	func() {
		fmt.Println("I'm driving!")
	}()
}

```
