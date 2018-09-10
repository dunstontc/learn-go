# Chapter 8: Goroutines & Channels 

<!-- TOC -->

- [8.1. Goroutines](#81-goroutines)
- [8.2. Example: Concurrent Clock Server](#82-example-concurrent-clock-server)
- [8.3. Example: Concurrent Echo Server](#83-example-concurrent-echo-server)
- [8.4. Channels 225 8.5. Looping in Parallel](#84-channels-225-85-looping-in-parallel)
- [8.6. Example: Concurrent Web Crawler](#86-example-concurrent-web-crawler)
- [8.7. Multiplexing with select](#87-multiplexing-with-select)
- [8.8. Example: Concurrent Directory Traversal](#88-example-concurrent-directory-traversal)
- [8.9. Cancellation](#89-cancellation)
- [8.10. Example: Chat Server](#810-example-chat-server)

<!-- /TOC -->

Concurrent programming, the expression of a program as a composition of several autonomous activities, has never been more important than it is today. Web servers handle requests for thousands of clients at once. Tablet and phone apps render animations in the user interface while simultaneously performing computation and network requests in the background. Even traditional batch problems—read some data, compute, write some output—use concurrency to hide the latency of I/O operations and to exploit a modern computer’s many processors, which every year grow in number but not in speed.

Go enables two styles of concurrent programming. This chapter presents goroutines and channels, which support *communicating sequential processes* or *CSP*, a model of concurrency in which values are passed between independent activities (goroutines) but variables are for the most part confined to a single activity. Chapter 9 covers some aspects of the more traditional model of *shared memory multithreading*, which will be familiar if you’ve used threads in other mainstream languages. Chapter 9 also points out some important hazards and pitfalls of concurrent programming that we won’t delve into in this chapter.

Even though Go’s support for concurrency is one of its great strengths, reasoning about concurrent programs is inherently harder than about sequential ones, and intuitions acquired from sequential programming may at times lead us astray. If this is your first encounter with concurrency, we recommend spending a little extra time thinking about the examples in these two chapters.


## 8.1. Goroutines 

In Go, each concurrently executing activity is called a *goroutine*. Consider a program that has two functions, one that does some computation and one that writes some output, and assume that neither function calls the other. A sequential program may call one function and then call the other, but in a *concurrent* program with two or more goroutines, calls to *both* functions can be active at the same time. We’ll see such a program in a moment.

If you have used operating system threads or threads in other languages, then you can assume for now that a goroutine is similar to a thread, and you’ll be able to write correct programs. The differences between threads and goroutines are essentially quantitative, not qualitative, and will be described in Section 9.8.

When a program starts, its only goroutine is the one that calls the main function, so we call it the *main goroutine*. New goroutines are created by the `go` statement. Syntactically, a `go` statement is an ordinary function or method call prefixed by the keyword `go`. A `go` statement causes the function to be called in a newly created goroutine. The `go` statement itself completes immediately:
```go
    f()    // call f(); wait for it to return
    go f() // create a new goroutine that calls f(); don't wait
```
In the example below, the main goroutine computes the 45th Fibonacci number. Since it uses the terribly inefficient recursive algorithm, it runs for an appreciable time, during which we’d like to provide the user with a visual indication that the program is still running, by displaying an animated textual "spinner."
```go
// gopl.io/ch8/spinner
// Spinner displays an animation while computing the 45th Fibonacci number.
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
```
After several seconds of animation, the `fib(45)` call returns and the `main` function prints its result:
```
    Fibonacci(45) = 1134903170
```
The `main` function then returns. When this happens, all goroutines are abruptly terminated and the program exits. Other than by returning from `main` or exiting the program, there is no programmatic way for one goroutine to stop another, but as we will see later, there are ways to communicate with a goroutine to request that it stop itself.

Notice how the program is expressed as the composition of two autonomous activities, spinning and Fibonacci computation. Each is written as a separate function but both make progress concurrently.


## 8.2. Example: Concurrent Clock Server 
## 8.3. Example: Concurrent Echo Server 
## 8.4. Channels 225 8.5. Looping in Parallel 
## 8.6. Example: Concurrent Web Crawler 
## 8.7. Multiplexing with select 
## 8.8. Example: Concurrent Directory Traversal 
## 8.9. Cancellation 
## 8.10. Example: Chat Server 
