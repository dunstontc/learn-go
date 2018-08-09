# Blank Identifier
- you must use everything you put in your code
- if you declare a variable, you must use it

## the blank identifier
- `_`
- allows you to tell the compiler you aren’t using something

### example

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		log.Fatal(err)
	}

	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", page)
}
```

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("https://jsonplaceholder.typicode.com/todos")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
```
