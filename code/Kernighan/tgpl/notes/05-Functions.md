# Chapter 5: Functions

<!-- TOC -->

- [5.1. Function Declarations](#51-function-declarations)
- [5.2. Recursion](#52-recursion)
- [5.3. Multiple Return Values](#53-multiple-return-values)
- [5.4. Errors](#54-errors)
  - [5.4.1. Error-Handling Strategies](#541-error-handling-strategies)
  - [5.4.2. End of File (EOF)](#542-end-of-file-eof)
- [5.5. Function Values](#55-function-values)
- [5.6. Anonymous Functions](#56-anonymous-functions)
  - [5.6.1. Caveat: Capturing Iteration Variables](#561-caveat-capturing-iteration-variables)
- [5.7. Variadic Functions](#57-variadic-functions)
- [5.8. Deferred Function Calls](#58-deferred-function-calls)
- [5.9. Panic](#59-panic)
- [5.10. Recover](#510-recover)

<!-- /TOC -->

A function lets us wrap up a sequence of statements as a unit that can be called from elsewhere in a program, perhaps multiple times. Functions make it possible to break a big job into smaller pieces that might well be written by different people separated by both time and space. A function hides its implementation details from its users. For all of these reasons, functions are a critical part of any programming language.

We’ve seen many functions already. Now let’s take time for a more thorough discussion. The running example of this chapter is a web crawler, that is, the component of a web search engine responsible for fetching web pages, discovering the links within them, fetching the pages identified by those links, and so on. A web crawler gives us ample opportunity to explore recursion, anonymous functions, error handling, and aspects of functions that are unique to Go.

## 5.1. Function Declarations 

A function declaration has a name, a list of parameters, an optional list of results, and a body:
```go
  func name(parameters) (results) { 
      body
  }
```
The parameter list specifies the names and types of the function’s *parameters*, which are the local variables whose values or *arguments* are supplied by the caller. The result list specifies the types of the values that the function returns. If the function returns one unnamed result or no results at all, parentheses are optional and usually omitted. Leaving off the result list entirely declares a function that does not return any value and is called only for its effects. In the hypot function,
```go
  func hypot(x, y float64) float64 {
      return math.Sqrt(x*x + y*y)
  }

  fmt.Println(hypot(3, 4)) // "5"
```
`x` and `y` are parameters in the declaration, `3` and `4` are arguments of the call, and the function returns a `float64` value.

Like parameters, results may be named. In that case, each name declares a local variable initialized to the zero value for its type.

A function that has a result list must end with a `return` statement unless execution clearly cannot reach the end of the function, perhaps because the function ends with a call to `panic` or an infinite `for` loop with no `break`.

As we saw with `hypot`, a sequence of parameters or results of the same type can be factored so that the type itself is written only once. These two declarations are equivalent:
```go
  func f(i, j, k int, s, t string)                { /* ... */ }
  func f(i int, j int, k int, s string, t string) { /* ... */ }
```
Here are four ways to declare a function with two parameters and one result, all of type `int`. The blank identifier can be used to emphasize that a parameter is unused.
```go
  func add(x int, y int) int { return x + y } 
  func sub(x, y int) (z int) { z = x - y; return } 
  func first(x int, _ int) int { return x }
  func zero(int, int) int

  fmt.Printf("%T\n", add)   // "func(int, int) int"
  fmt.Printf("%T\n", sub)   // "func(int, int) int"
  fmt.Printf("%T\n", first) // "func(int, int) int"
  fmt.Printf("%T\n", zero)  // "func(int, int) int"
```

The type of a function is sometimes called its *signature*. Two functions have the same type or signature if they have the same sequence of parameter types and the same sequence of result types. The names of parameters and results don’t affect the type, nor does whether or not they were declared using the factored form.

Every function call must provide an argument for each parameter, in the order in which the parameters were declared. Go has no concept of default parameter values, nor any way to specify arguments by name, so the names of parameters and results don’t matter to the caller except as documentation.

Parameters are local variables within the body of the function, with their initial values set to the arguments supplied by the caller. Function parameters and named results are variables in the same lexical block as the function’s outermost local variables.

Arguments are *passed by value*, so the function receives a copy of each argument; modifications to the copy do not affect the caller. However, if the argument contains some kind of reference, like a pointer, slice, map, function, or channel, then the caller may be affected by any modifications the function makes to variables *indirectly* referred to by the argument.

You may occasionally encounter a function declaration without a body, indicating that the function is implemented in a language other than Go. Such a declaration defines the function signature:
```go
  package math

  func Sin(x float64) float64 // implemented in assembly language
```


## 5.2. Recursion 

Functions may be *recursive*, that is, they may call themselves, either directly or indirectly. Recursion is a powerful technique for many problems, and of course it’s essential for processing recursive data structures. In Section 4.4, we used recursion over a tree to implement a simple insertion sort. In this section, we’ll use it again for processing HTML documents.

The example program below uses a non-standard package, `golang.org/x/net/html`, which provides an HTML parser. The `golang.org/x/...` repositories hold packages designed and maintained by the Go team for applications such as networking, internationalized text processing, mobile platforms, image manipulation, cryptography, and developer tools. These packages are not in the standard library because they’re still under development or because they’re rarely needed by the majority of Go programmers.

The parts of the `golang.org/x/net/html` API that we’ll need are shown below. The function `html.Parse` reads a sequence of bytes, parses them, and returns the root of the HTML document tree, which is an `html.Node.` HTML has several kinds of nodes (text, comments, and so on) but here we are concerned only with *element* nodes of the form `<name key='value'>`.
```go
// golang.org/x/net/html
  package html

  type Node struct {
      Type                    NodeType
      Data                    string
      Attr                    []Attribute
      FirstChild, NextSibling *Node
  }

  type NodeType int32

  const (
      ErrorNode NodeType = iota
      TextNode
      DocumentNode
      ElementNode
      CommentNode
      DoctypeNode
  )

  type Attribute struct {
      Key, Val string
  }

  func Parse(r io.Reader) (*Node, error)
```

The `main` function parses the standard input as HTML, extracts the links using a recursive `visit` function, and prints each discovered link:
```go
// gopl.io/ch5/findlinks1
// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
```
The `visit` function traverses an HTML node tree, extracts the link from the `href` attribute of each *anchor* element `<a href='...'>`, appends the links to a slice of strings, and returns the resulting slice:
```go
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
```
To descend the tree for a node `n`, visit recursively calls itself for each of `n`'s children, which are held in the `FirstChild` linked list.  

Let’s run `findlinks` on the Go home page, piping the output of `fetch` (§1.5) to the input of `findlinks`. We've edited the output slightly for brevity.

```
  $ go build gopl.io/ch1/fetch
  $ go build gopl.io/ch5/findlinks1
  $ ./fetch https://golang.org | ./findlinks1
  #
  /doc/
  /pkg/
  /help/
  /blog/
  http://play.golang.org/
  //tour.golang.org/
  https://golang.org/dl/
  //blog.golang.org/
  /LICENSE
  /doc/tos.html
  http://www.google.com/intl/en/policies/privacy/
```

Notice the variety of forms of links that appear in the page. Later we’ll see how to resolve them relative to the base URL, `https://golang.org`, to make absolute URLs.

The next program uses recursion over the HTML node tree to print the structure of the tree in outline. As it encounters each element, it pushes the element’s tag onto a stack, then prints the stack.
```go
// gopl.io/ch5/outline
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
```
Note one subtlety: although `outline` "pushes" an element on stack, there is no corresponding pop. When `outline` calls itself recursively, the callee receives a copy of `stack`. Although the callee may append elements to this slice, modifying its underlying array and perhaps even allocating a new array, it doesn’t modify the initial elements that are visible to the caller, so when the function returns, the caller’s `stack` is as it was before the call.

Here’s the outline of `https://golang.org`, again edited for brevity:
```
  $ go build gopl.io/ch5/outline
  $ ./fetch https://golang.org | ./outline
  [html]
  [html head]
  [html head meta]
  [html head title]
  [html head link]
  [html body]
  [html body div]
  [html body div]
  [html body div div]
  [html body div div form]
  [html body div div form div]
  [html body div div form div a]
  ...
```

As you can see by experimenting with `outline`, most HTML documents can be processed with only a few levels of recursion, but it’s not hard to construct pathological web pages that require extremely deep recursion.

Many programming language implementations use a fixed-size function call stack; sizes from 64KB to 2MB are typical. Fixed-size stacks impose a limit on the depth of recursion, so one must be careful to avoid a *stack overflow* when traversing large data structures recursively; fixed-size stacks may even pose a security risk. In contrast, typical Go implementations use variable-size stacks that start small and grow as needed up to a limit on the order of a gigabyte. This lets us use recursion safely and without worrying about overflow.

### Exercises
- **Exercise 5.1**: Change the `findlinks` program to traverse the `n.FirstChild` linked list using recursive calls to `visit` instead of a loop.
- **Exercise 5.2**: Write a function to populate a mapping from element names (`p`, `div`, `span`, and so on) to the number of elements with that name in an HTML document tree.
- **Exercise 5.3**: Write a function to print the contents of all text nodes in an HTML document tree. Do not descend into `<script>` or `<style>` elements, since their contents are not visible in a web browser.
- **Exercise 5.4**: Extend the `visit` function so that it extracts other kinds of links from the document, such as images, scripts, and style sheets.

## 5.3. Multiple Return Values 
## 5.4. Errors 
### 5.4.1. Error-Handling Strategies
### 5.4.2. End of File (EOF)
## 5.5. Function Values 
## 5.6. Anonymous Functions 
### 5.6.1. Caveat: Capturing Iteration Variables
## 5.7. Variadic Functions 
## 5.8. Deferred Function Calls 
## 5.9. Panic 
## 5.10. Recover



