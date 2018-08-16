package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fatih/color"
)

var wg sync.WaitGroup
var counter int64

func main() {
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	wg.Add(2)
	go incrementor(blue("Foo:"))
	go incrementor(yellow("Bar:"))
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	green := color.New(color.FgGreen).SprintFunc()
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, green("Counter:"), atomic.LoadInt64(&counter)) // access without race
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
