# The Go Programming Language

## Contents

0. Preface
    - Preface
    - The Origins of Go
    - The Go Project
    - Organization of the Book
    - Where to Find More Information Acknowledgments
1. Tutorial
    - 1.1. Hello, World
    - 1.2. Command-Line Arguments 
    - 1.3. Finding Duplicate Lines
    - 1.4. Animated GIFs
    - 1.5. Fetching a URL
    - 1.6. Fetching URLs Concurrently 
    - 1.7. A Web Server
    - 1.8. Loose Ends
2. Program Structure
    - 2.1. Names
    - 2.2. Declarations
     - 2.3.1 Short Variable Declarations
      - 2.3.2 Pointers
      - 2.3.3 The `new` Function
      - 2.3.4 Lifetime of Variables
      - 2.3.5 Tuple Assignment
      - 2.3.6 Assignability
    - 2.3. Variables
    - 2.4. Assignments
    - 2.5. Type Declarations 
    - 2.6. Packages and Files 
    - 2.7. Scope
3. Basic Data Types
    - 3.1. Integers 
    - 3.2. Floating-Point Numbers 
    - 3.3. Complex Numbers 
    - 3.4. Booleans 
    - 3.5. Strings 
      - 3.5.1. String Literals
      - 3.5.2. Unicode
      - 3.5.3. UTF-8
      - 3.5.4. Strings and Byte Slices
      - 3.5.5. Conversions between Strings and Numbers
    - 3.6. Constants
4. Composite Types
    - 4.1. Arrays 
    - 4.2. Slices 
      - 4.2.1. TheappendFunction
    - 4.3. Maps 
    - 4.4. Structs 
      - 4.4.1. Struct Literals
      - 4.4.2. Comparing Structs
      - 4.4.3. Struct Embedding and Anonymous Fields
    - 4.5. JSON 
    - 4.6. Text and HTML Templates
5. Functions
    - 5.1. Function Declarations 
    - 5.2. Recursion 
    - 5.3. Multiple Return Values 
    - 5.4. Errors 
      - 5.4.1. Error-Handling Strategies
      - 5.4.2. End of File (EOF)
    - 5.5. Function Values 
    - 5.6. Anonymous Functions 
      - 5.6.1. Caveat: Capturing Iteration Variables
    - 5.7. Variadic Functions 
    - 5.8. Deferred Function Calls 
    - 5.9. Panic 
    - 5.10. Recover
6. Methods 
    - 6.1. Method Declarations 
    - 6.2. Methods with a Pointer Receiver 
    - 6.3. Composing Types by Struct Embedding 
    - 6.4. Method Values and Expressions 
    - 6.5. Example: Bit Vector Type 
    - 6.6. Encapsulation
7. Interfaces
    - 7.1. Interfaces as Contracts 
    - 7.2. Interface Types 
    - 7.3. Interface Satisfaction 
    - 7.4. Parsing Flags with flag.Value 
    - 7.5. Interface Values
    - 7.6. Sorting with sort.Interface 
    - 7.7. The http.Handler Interface 
    - 7.8. The error Interface 
    - 7.9. Example: Expression Evaluator 
    - 7.10. Type Assertions 
    - 7.11. Discriminating Errors with Type Assertions 
    - 7.12. Querying Behaviors with Interface Type Assertions 
    - 7.13. Type Switches 
    - 7.14. Example: Token-Based XML Decoding 
    - 7.15. A Few Words of Advice
8. Goroutines and Channels
    - 8.1. Goroutines 
    - 8.2. Example: Concurrent Clock Server 
    - 8.3. Example: Concurrent Echo Server 
    - 8.4. Channels 225 8.5. Looping in Parallel 
    - 8.6. Example: Concurrent Web Crawler 
    - 8.7. Multiplexing with select 
    - 8.8. Example: Concurrent Directory Traversal 
    - 8.9. Cancellation 
    - 8.10. Example: Chat Server 
9. Concurrency with Shared Variables
    - 9.1. Race Conditions 
    - 9.2. Mutual Exclusion: sync.Mutex 
    - 9.3. Read/Write Mutexes: sync.RWMutex 
    - 9.4. Memory Synchronization 
    - 9.5. Lazy Initialization: sync.Once 
    - 9.6. The Race Detector 
    - 9.7. Example: Concurrent Non-Blocking Cache 
    - 9.8. Goroutines and Threads
10. Packages and the Go Tool
    - 10.1. Introduction
    - 10.2. Import Paths
    - 10.3. The Package Declaration 
    - 10.4. Import Declarations 
    - 10.5. Blank Imports
    - 10.6. Packages and Naming 
    - 10.7. The Go Tool
11. Testing
    - 11.1. The go test Tool 
    - 11.2. Test Functions 
    - 11.3. Coverage 
    - 11.4. Benchmark Functions 
    - 11.5. Profiling 
    - 11.6. Example Functions 
12. Reflection
    - 12.1. Why Reflection? 
    - 12.2. reflect.Type and reflect.Value 
    - 12.3. Display, a Recursive Value Printer 
    - 12.4. Example: Encoding S-Expressions 
    - 12.5. Setting Variables with reflect.Value 
    - 12.6. Example: Decoding S-Expressions 
    - 12.7. Accessing Struct Field Tags 
    - 12.8. Displaying the Methods of a Type 
    - 12.9. A Word of Caution 
13. Low-Level Programming
    - 13.1. unsafe.Sizeof, Alignof, and Offsetof
    - 13.2. unsafe.Pointer
    - 13.3. Example: Deep Equivalence 
    - 13.4. Calling C Code with cgo
    - 13.5. Another Word of Caution
