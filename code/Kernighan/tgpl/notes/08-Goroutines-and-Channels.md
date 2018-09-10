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

Networking is a natural domain in which to use concurrency since servers typically handle many connections from their clients at once, each client being essentially independent of the others. In this section, we’ll introduce the `net` package, which provides the components for building networked client and server programs that communicate over TCP, UDP, or Unix domain sockets. The `net/http` package we’ve been using since Chapter 1 is built on top of functions from the `net` package.

Our first example is a sequential clock server that writes the current time to the client once per second:
```go
// gopl.io/ch8/clock1
// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
```
The `Listen` function creates a `net.Listener`, an object that listens for incoming connections on a network port, in this case TCP port `localhost:8000`. The listener’s `Accept` method blocks until an incoming connection request is made, then returns a `net.Conn` object representing the connection.

The `handleConn` function handles one complete client connection. In a loop, it writes the current time, `time.Now()`, to the client. Since `net.Conn` satisfies the `io.Writer` interface, we can write directly to it. The loop ends when the write fails, most likely because the client has disconnected, at which point `handleConn` closes its side of the connection using a deferred call to `Close` and goes back to waiting for another connection request.

The `time.Time.Format` method provides a way to format date and time information by example. Its argument is a template indicating how to format a reference time, specifically `Mon Jan 2 03:04:05PM 2006 UTC-0700`. The reference time has eight components (dayofthe week, month, day of the month, and so on). Any collection of them can appear in the `Format` string in any order and in a number of formats; the selected components of the date and time will be displayed in the selected formats. Here we are just using the hour, minute, and second of the time. The `time` package defines templates for many standard time formats, such as `time.RFC1123`. The same mechanism is used in reverse when parsing a time using `time.Parse`.

To connect to the server, we’ll need a client program such as `nc` ("netcat"), a standard utility program for manipulating network connections:
```
    $ go build gopl.io/ch8/clock1
    $ ./clock1 &
    $ nc localhost 8000
    13:58:54
    13:58:55
    13:58:56
    13:58:57
    ^C
```
The client displays the time sent by the server each second until we interrupt the client with Control-C, which on Unix systems is echoed as `^C` by the shell. If `nc` or `netcat` is not installed on your system, you can use `telnet` or this simple Go version of `netcat` that uses `net.Dial` to connect to a TCP server:
```go
// gopl.io/ch8/netcat1
// Netcat1 is a read-only TCP client.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
```
This program reads data from the connection and writes it to the standard output until an end-of-file condition or an error occurs. The `mustCopy` function is a utility used in several examples in this section. Let’s run two clients at the same time on different terminals, one shown to the left and one to the right:
```
$ go build gopl.io/ch8/netcat1
$ ./netcat1
13:58:54                           $ ./netcat1
13:58:55
13:58:56 
^C
                                   13:58:57
                                   13:58:58
                                   13:58:59
                                   ^C
$ killall clock1
```
The `killall` command is a Unix utility that kills all processes with the given name.

The second client must wait until the first client is finished because the server is *sequential*; it deals with only one client at a time. Just one small change is needed to make the server concurrent: adding the `go` keyword to the call to `handleConn` causes each call to run in its own goroutine.
```go
// gopl.io/ch8/clock2

```
Now, multiple clients can receive the time at once:
```
$ go build gopl.io/ch8/clock2
$ ./clock2 &
$ go build gopl.io/ch8/netcat1
$ ./netcat1
14:02:54                          $ ./netcat1
14:02:55                          14:02:55
14:02:56                          14:02:56
14:02:57                          ^C
14:02:58
14:02:59                          $ ./netcat1
14:03:00                          14:03:00
14:03:01                          14:03:01
^C                                14:03:02
                                  ^C
$ killall clock2
```

### Exercises
- **Exercise 8.1**: Modify `clock2` to accept a port number, and write a program, `clockwall`, that acts as a client of several clock servers at once, reading the times from each one and displaying the results in a table, akin to the wall of clocks seen in some business offices. If you have access to geographically distributed computers, run instances remotely; otherwise run local instances on different ports with fake time zones.
```
    $ TZ=US/Eastern    ./clock2 -port 8010 &
    $ TZ=Asia/Tokyo    ./clock2 -port 8020 &
    $ TZ=Europe/London ./clock2 -port 8030 &
    $ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
```
- **Exercise 8.2**: Implement a concurrent File Transfer Protocol (FTP) server. The server should interpret commands from each client such as `cd` to change directory, `ls` to list a directory, `get` to send the contents of a file, and `close` to close the connection. You can use the standard `ftp` command as the client, or write your own.


## 8.3. Example: Concurrent Echo Server 

The clock server used one goroutine per connection. In this section, we’ll build an echo server that uses multiple goroutines per connection. Most echo servers merely write whatever they read, which can be done with this trivial version of `handleConn`:
```go
    func handleConn(c net.Conn) {
        io.Copy(c, c) // NOTE: ignoring errors
        c.Close()
    }
```
A more interesting echo server might simulate the reverberations of a real echo, with the response loud at first (`"HELLO!"`), then moderate (`"Hello!"`) after a delay, then quiet (`"hello!"`) before fading to nothing, as in this version of handleConn:
```go
// gopl.io/ch8/reverb1
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}
```
We’ll need to upgrade our client program so that it sends terminal input to the server while also copying the server response to the output, which presents another opportunity to use concurrency:
```go
// gopl.io/ch8/netcat2
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
```
While the main goroutine reads the standard input and sends it to the server, a second goroutine reads and prints the server’s response. When the main goroutine encounters the end of the input, for example, after the user types Control-D (`^D`) at the terminal (or the equivalent Control-Z on Microsoft Windows), the program stops, even if the other goroutine still has work to do. (We’ll see how to make the program wait for both sides to finish once we’ve introduced channels in Section 8.4.1.)

In the session below, the client’s input is left-aligned and the server’s responses are indented. 
The client shouts at the echo server three times:
```
    $ go build gopl.io/ch8/reverb1
    $ ./reverb1 &
    $ go build gopl.io/ch8/netcat2
    $ ./netcat2
    Hello?
        HELLO?
        Hello?
        hello?
    Is there anybody there?
        IS THERE ANYBODY THERE?
    Yooo-hooo!
        Is there anybody there?
        is there anybody there?
        YOOO-HOOO!
        Yooo-hooo!
        yooo-hooo!
    ^D
    $ killall reverb1
```
Notice that the third shout from the client is not dealt with until the second shout has petered out, which is not very realistic. A real echo would consist of the *composition* of the three independent shouts. To simulate it, we’ll need more goroutines. Again, all we need to do is add the `go` keyword, this time to the call to `echo`:
```go
// gopl.io/ch8/reverb2
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}
```
The arguments to the function started by `go` are evaluated when the `go` statement itself is executed; thus `input.Text()` is evaluated in the main goroutine.

Now the echoes are concurrent and overlap in time:
```
  $ go build gopl.io/ch8/reverb2
  $ ./reverb2 &
  $ ./netcat2
  Is there anybody there?
      IS THERE ANYBODY THERE?
  Yooo-hooo!
      Is there anybody there?
      YOOO-HOOO!
      is there anybody there?
      Yooo-hooo!
      yooo-hooo!
    ^D
    $ killall reverb2
```
All that was required to make the server use concurrency, not just to handle connections from multiple clients but even within a single connection, was the insertion of two `go` keywords.

However in adding these keywords, we had to consider carefully that it is safe to call methods of `net.Conn` concurrently, which is not true for most types. We’ll discuss the crucial concept of *concurrency safety* in the next chapter.


## 8.4. Channels 
### 8.4.1. Unbuffered Channels
### 8.4.2. Pipelines
### 8.4.3. Unidirectional Channel Types
### 8.4.4. Buffered Channels
## 8.5. Looping in Parallel 
## 8.6. Example: Concurrent Web Crawler 
## 8.7. Multiplexing with select 
## 8.8. Example: Concurrent Directory Traversal 
## 8.9. Cancellation 
## 8.10. Example: Chat Server 
