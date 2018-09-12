# Chapter 11: Testing

<!-- TOC -->

- [11.1. The go test Tool](#111-the-go-test-tool)
- [11.2. Test Functions](#112-test-functions)
  - [11.2.1. Randomized Testing](#1121-randomized-testing)
  - [11.2.2. Testing a Command](#1122-testing-a-command)
  - [11.2.3. White-Box Testing](#1123-white-box-testing)
  - [11.2.4. External Test Packages](#1124-external-test-packages)
  - [11.2.5. Writing Effective Tests](#1125-writing-effective-tests)
  - [11.2.6. Avoiding Brittle Tests](#1126-avoiding-brittle-tests)
- [11.3. Coverage](#113-coverage)
- [11.4. Benchmark Functions](#114-benchmark-functions)
- [11.5. Profiling](#115-profiling)
- [11.6. Example Functions](#116-example-functions)

<!-- /TOC -->

Maurice Wilkes, the developer of EDSAC, the first stored-program computer, had a startling insight while climbing the stairs of his laboratory in 1949. In *Memoirs of a Computer Pioneer*, he recalled, 
> "The realization came over me with full force that a good part of the remainder of my life was going to be spent in finding errors in my own programs." 
Surely every programmer of a stored-program computer since then can sympathize with Wilkes, though perhaps not without some bemusement at his naïveté about the difficulties of software construction.

Programs today are far larger and more complex than in Wilkes’s time, of course, and a great deal of effort has been spent on techniques to make this complexity manageable. Two techniques in particular stand out for their effectiveness. The first is routine peer review of programs before they are deployed. The second, the subject of this chapter, is testing.

Testing, by which we implicitly mean *automated testing*, is the practice of writing small programs that check that the code under test (the production code) behaves as expected for certain inputs, which are usually either carefully chosen to exercise certain features or randomized to ensure broad coverage.

The field of software testing is enormous. The task of testing occupies all programmers some of the time and some programmers all of the time. The literature on testing includes thousands of printed books and millions of words of blog posts. In every mainstream programming language, there are dozens of software packages intended for test construction, some with a great deal of theory, and the field seems to attract more than a few prophets with cult-like followings. It is almost enough to convince programmers that to write effective tests they must acquire a whole new set of skills.

Go’s approach to testing can seem rather low-tech in comparison. It relies on one command, go test, and a set of conventions for writing test functions that go test can run. The comparatively lightweight mechanism is effective for pure testing, and it extends naturally to benchmarks and systematic examples for documentation.

In practice, writing test code is not much different from writing the original program itself. We write short functions that focus on one part of the task. We have to be careful of boundary conditions, think about data structures, and reason about what results a computation should produce from suitable inputs. But this is the same process as writing ordinary Go code; it needn’t require new notations, conventions, and tools.

## 11.1. The go test Tool 
## 11.2. Test Functions 
### 11.2.1. Randomized Testing
### 11.2.2. Testing a Command
### 11.2.3. White-Box Testing
### 11.2.4. External Test Packages
### 11.2.5. Writing Effective Tests
### 11.2.6. Avoiding Brittle Tests
## 11.3. Coverage 
## 11.4. Benchmark Functions 
## 11.5. Profiling 
## 11.6. Example Functions 
