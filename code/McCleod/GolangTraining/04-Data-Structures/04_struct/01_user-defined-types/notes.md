## Go is Object Oriented

### Encapsulation  
- state (*fields*)  
- behavior (*methods*)  
- exported / un-exported  

### Reusability
- inheritence (*embedded types*)

### Polymorphism
- interfaces

### Overriding
- *promotion*


## Traditional OOP

### Classes
- data structure describing a type of object
- you can then create "instances"/"objects" from the class/blue-print
- classes hold both:
  - state / data / fields
  - behavior / methods
  - Public / private

### Inheritence

## In Go:
- you don't create classes, you create a *type*
- you don't instantiate, you create a value of a type

---

user defined types - we declare a new type, foo
the underlying type of foo: int

conversion:int(myAge)
converting type foo to type int

THIS CODE IS ONLY FOR EXAMPLE
IT IS A BAD PRACTICE TO ALIAS TYPES
one exception: if you need to attach methods to a type, see the time package for an example of this

```go
//  godoc.org/time
type Duration int64
```

Duration has methods attached to it
