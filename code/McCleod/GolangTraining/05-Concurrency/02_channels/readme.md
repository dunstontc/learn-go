# Channels

## Concepts
- Unbuffered channels block
- You can return channels
- You can take channels as arguments
- Patterns
  - Pipelines
  - Fan Out, Fan In
- Channel Direction




## Range Clause


## N to 1
  - Many functions writing to the same channel
  - (No `sync.WaitGroup().ADD()` inside goroutines)
```go
func main() {

	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
```


## 1 to N


## Pass/Return


## Channel Direction

> The optional `<-` operator specifies the channel *direction*, *send* or *receive*.   
> If no direction is given, the channel is *bidirectional*. A channel may be constrained only to send or only to receive by *conversion* or *assignment*.
