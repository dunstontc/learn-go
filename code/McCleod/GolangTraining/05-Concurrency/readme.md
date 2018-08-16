# Concurrency


## parallelism
But when people hear the word *concurrency* they often think of *parallelism*, a related but quite distinct concept. In programming, *concurrency* is the *composition* of independently executing processes, while parallelism is the simultaneous *execution* of (possibly related) computations. *Concurrency* is about dealing with lots of things at once. *Parallelism* is about doing lots of things at once.

### without concurrency
```go
func main() {
	foo()
	bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
```


## waitgroups
A *WaitGroup* waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

### with concurrency
```go
func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
```


## race conditions
*Race conditions* are among the most insidious and elusive programming errors. They typically cause erratic and mysterious failures, often long after the code has been deployed to production. While Go's concurrency mechanisms make it easy to write clean concurrent code, they don't prevent race conditions. Care, diligence, and testing are required. And tools can help.

### with waitgroups
```go
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
```

## mutex
A *Mutex* is a *mutual exclusion lock*. Mutexes can be created as part of other structures; the zero value for a Mutex is an unlocked mutex.

## atomicity
- Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.
- These functions require great care to be used correctly. Except for special, low-level applications, synchronization is better done with channels or the facilities of the sync package. Share memory by communicating; don't communicate by sharing memory.

## channels
