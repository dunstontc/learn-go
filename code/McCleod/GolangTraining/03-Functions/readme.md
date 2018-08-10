# Functions

- Parameters vs. Arguments
  - Params
    - multiple *variadic* params
  - Args
    - multiple *variadic* args
- Returns
  - multiple returns
  - named returns - yuck!
<!-- - review
  - func expressions
  - closure  -->
- Callbacks
- Recursion
- Defer
- Pass by Value
  - reference types
- Immediately Invoked Function Expressions (iifes)


## Parameters v. Arguments

> You define with params and call with args.


## Returns
```go
func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
```

### Multiple Returns
```go
func main() {
	fmt.Println(greet("Jane ", "Doe "))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
```

### Named Returns
```go
func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
```


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
