# Boolean Expressions

- expressions vs statements
- Bool Types
  - `true`
  - `false`
- Operators  
  - not   
    - `!`
  - or   
    - `||`  
  - and  
    - `&&`  


## True-False
```go
func main() {

	if true {
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run")
	}
}
```

## Operators

### Not (Bang)
```go
func main() {

	if !true {
		fmt.Println("This did not run")
	}

	if !false {
		fmt.Println("This ran")
	}

}
```

### Or 
```go
func main() {

	if true || false {
		fmt.Println("This ran")
	}
}
```

### And 
```go
func main() {

	if true && false {
		fmt.Println("This did not run")
	}

}
```
