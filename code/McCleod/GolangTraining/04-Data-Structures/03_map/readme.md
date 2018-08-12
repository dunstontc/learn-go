# Maps

## Making Maps

```go
func main() {

	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting) // map[Tim:Good morning. Jenny:Bonjour.]
}
```

```go
func main() {

	myGreeting := make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting) // map[Tim:Good morning. Jenny:Bonjour.]
}
```

```go
func main() {

	myGreeting := map[string]string{}
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting) // map[Tim:Good morning. Jenny:Bonjour.]
}
```

```go
func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	fmt.Println(myGreeting["Jenny"]) // Bonjour!
}
```
