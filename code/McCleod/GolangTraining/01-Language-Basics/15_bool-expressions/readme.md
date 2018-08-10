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


## Operators

### Not

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
