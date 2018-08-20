for your perusal:
https://golang.org/ref/spec#Composite_literals


```go
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"","","",},
				course{"","","",},
				course{"","","",},
				course{"","","",},
				course{"","","",},
				course{"","","",},
				course{"","","",},
				course{"","","",},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"","","",},
				course{"","","",},
				course{"","","",},
			},
		},
	}
```
