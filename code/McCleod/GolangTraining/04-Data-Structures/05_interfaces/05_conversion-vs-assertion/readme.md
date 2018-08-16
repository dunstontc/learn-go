# Conversion v. Assertion



## Conversion
- Widening v Narrowing Conversion

### int to float
```go
func main() {
	var x = 12
	var y = 12.1230123
	fmt.Println(y + float64(x)) // 24.1230123
	// conversion: int to float64
}
```

### float to int
```go
func main() {
	var x = 12
	var y = 12.1230123
	fmt.Println(int(y) + x) // 24
	// conversion: float64 to int
}
```

### rune to string
```go
func main() {
	var x rune = 'a' // rune is an alias for int32; normally omitted in this statement
	var y int32 = 'b'
	fmt.Println(x)         // 97
	fmt.Println(y)         // 98
	fmt.Println(string(x)) // a
	fmt.Println(string(y)) // b
	// conversion: rune to string
}
```

### rune to slice of bytes to string
```go
func main() {
	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'})) // hello
	// conversion: []bytes to string
}
```

### string to slice of bytes
```go
func main() {
	fmt.Println([]byte("hello")) // [104 101 108 108 111]
	// conversion: string to []bytes
}
```

### strconv
- `Atoi()` - (ascii) string to int
- `strconv.Itoa()` - int to string (ascii)

#### Atoi
```go
import (
	"fmt"
	"strconv"
)

func main() {
	var x = "12"
	var y = 6
	z, _ := strconv.Atoi(x)
  fmt.Println(y + z) // 18
}
```

#### Itoa
```go
import (
	"fmt"
	"strconv"
)

func main() {
  x := 12
	y := "I have this many: " + strconv.Itoa(x)
	fmt.Println(y) // I have this many: 12
}
```

#### Parseint
```go
import (
	"fmt"
	"strconv"
)

func main() {

	//	ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, i, u)

	//	FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
	w := strconv.FormatBool(true)
	x := strconv.FormatFloat(3.1415, 'E', -1, 64)
	y := strconv.FormatInt(-42, 16)
	z := strconv.FormatUint(42, 16)

	fmt.Println(w, x, y, z)
}
```


## Assertion
- Assertion is just for interfaces

```go
func main() {
	var name interface{} = "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str) // string
	} else {
		fmt.Printf("value is not a string\n")
	}
}
```

```go
func main() {
	var name interface{} = 7
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n") // value is not a string
	}
}
```

```go
func main() {
	var val interface{} = 7
	fmt.Printf("%T\n", val) // int
}
```

```go
func main() {
	var val interface{} = 7
	fmt.Println(val + 6)
	// ./main.go:7:18: invalid operation: val + 6 (mismatched types interface {} and int)
}
```

```go
func main() {
	var val interface{} = 7
	fmt.Println(val.(int) + 6) // 13
}
```

```go
func main() {
	rem := 7.24
	fmt.Printf("%T\n", rem)      // float64
	fmt.Printf("%T\n", int(rem)) // int
}
```

### Casting vs Assertion
```go
int(num)  // casting
num.(int) // assertion
```
