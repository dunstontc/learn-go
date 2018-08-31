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

Go enables two styles of concurrent programming. This chapter presents goroutines and channels, which support *communicating sequential processes or CSP*, a model of concurrency in which values are passed between independent activities (goroutines) but variables are for the most part confined to a single activity. Chapter 9 covers some aspects of the more traditional model of *shared memory multithreading*, which will be familiar if you’ve used threads in other mainstream languages. Chapter 9 also points out some important hazards and pitfalls of concurrent programming that we won’t delve into in this chapter.

Even though Go’s support for concurrency is one of its great strengths, reasoning about concurrent programs is inherently harder than about sequential ones, and intuitions acquired from sequential programming may at times lead us astray. If this is your first encounter with concurrency, we recommend spending a little extra time thinking about the examples in these two chapters.

## 8.1. Goroutines 
## 8.2. Example: Concurrent Clock Server 
## 8.3. Example: Concurrent Echo Server 
## 8.4. Channels 225 8.5. Looping in Parallel 
## 8.6. Example: Concurrent Web Crawler 
## 8.7. Multiplexing with select 
## 8.8. Example: Concurrent Directory Traversal 
## 8.9. Cancellation 
## 8.10. Example: Chat Server 
