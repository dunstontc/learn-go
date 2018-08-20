# Passing data to templates

These files provide you with examples of passing various data types to templates.

## Slices
```go
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
```

```html
<ul>
    {{range .}}
      <li> {{.}} </li>
    {{end}}
</ul>
```

## Maps
```go
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
```

```html
<ul>
  {{range .}}
    <li> {{.}} </li>
  {{end}}
</ul>
```

## Structs

```go
var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
```

```html
<ul>
    <li> {{.Name}} - {{.Motto}} </li>
</ul>
```

## Slice of Structs
```go
var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{Name: "Buddha", Motto: "The belief of no beliefs"}
	gandhi := sage{Name: "Gandhi", Motto: "Be the change"}
	mlk := sage{Name: "Martin Luther King", Motto: "Hatred never ceases with hatred but with love alone is healed."}
  jesus := sage{Name: "Jesus", Motto: "Love all"}
	muhammad := sage{Name: "Muhammad" Motto: "To overcome evil with good is good, to resist evil by evil is evil."}
  
	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
```

```html
<ul>
  {{range .}}
    <li> {{.Name}} - {{.Motto}} </li>
  {{end}}
</ul>
```

## Struct of Slice of Struct

```go
var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
```

```html
<ul>
    {{range .Wisdom}}
        <li> {{.Name}} - {{.Motto}} </li>
    {{end}}
</ul>
<ul>
    {{range .Transport}}
        <li> {{.Manufacturer}} - {{.Model}} - {{.Doors}} </li>
    {{end}}
</ul>
```
