# Channel 9: Concurrency with Shared Variables

<!-- TOC -->

- [9.1. Race Conditions](#91-race-conditions)
- [9.2. Mutual Exclusion: sync.Mutex](#92-mutual-exclusion-syncmutex)
- [9.3. Read/Write Mutexes: sync.RWMutex](#93-readwrite-mutexes-syncrwmutex)
- [9.4. Memory Synchronization](#94-memory-synchronization)
- [9.5. Lazy Initialization: sync.Once](#95-lazy-initialization-synconce)
- [9.6. The Race Detector](#96-the-race-detector)
- [9.7. Example: Concurrent Non-Blocking Cache](#97-example-concurrent-non-blocking-cache)
- [9.8. Goroutines and Threads](#98-goroutines-and-threads)

<!-- /TOC -->

In the previous chapter, we presented several programs that use goroutines and channels to express concurrency in a direct and natural way. However, in doing so, we glossed over a number of important and subtle issues that programmers must bear in mind when writing concurrent code.

In this chapter, we’ll take a closer look at the mechanics of concurrency. In particular, we’ll point out some of the problems associated with sharing variables among multiple goroutines, the analytical techniques for recognizing those problems, and the patterns for solving them. Finally, we’ll explain some of the technical differences between goroutines and operating system threads.


## 9.1. Race Conditions 
## 9.2. Mutual Exclusion: sync.Mutex 
## 9.3. Read/Write Mutexes: sync.RWMutex 
## 9.4. Memory Synchronization 
## 9.5. Lazy Initialization: sync.Once 
## 9.6. The Race Detector 
## 9.7. Example: Concurrent Non-Blocking Cache 
## 9.8. Goroutines and Threads
### 9.8.1
### 9.8.2
### 9.8.3
### 9.8.4
