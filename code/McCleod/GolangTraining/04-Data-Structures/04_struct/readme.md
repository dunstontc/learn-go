# Structs


## Go is Object Oriented
- **Encapsulation**  
  - state (*fields*)  
  - behavior (*methods*)  
  - exported / un-exported  
- **Reusability**
  - inheritence (*embedded types*)
- **Polymorphism**
  - interfaces
- **Overriding**
  - *promotion*

### Classes (traditional OOP)
- data structure describing a type of object
- you can then create "instances"/"objects" from the class/blue-print
- classes hold both:
  - state / data / fields
  - behavior / methods
  - Public / private

### In Go:
- you don't create *classes*, you create a `type`
- you don't *instantiate*, you create a value of a `type`
- here is how we talk about structs in Go:
  - user defined type
  - we *declare* the type
  - the type has *fields*
  - the type can also have *"tags"*
  - the type has an underlying type
    - in this case, the underlying type is struct
 - we declare variables of the type
 - we *initialize* those variables
  - initialize with a specific value or initiliaze to the zero value
  - a struct is a composite type


## User Defined Types


## Struct Fields, Values, & Initialization
```go
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
```


## Methods
```go
type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
```


## Embedded Types
```go
type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)
}
```


## Promotion
- The outermost fields & methods are defaulted to; use dot notation to dig into embedded data.


## Struct Pointer
```go
type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"James", 20}

	fmt.Println(p1)        // &{James 20}
	fmt.Printf("%T\n", p1) // *main.person
	fmt.Println(p1.name)   // James
	fmt.Println(p1.age)    // 20
}
```


## JSON

- Marshal/Unmarshal 
  - Strings
- Encode/Decode
  - Streams
  - Encode (Writer)
  - Decode (Reader)
