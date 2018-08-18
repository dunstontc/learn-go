// Package main provides Object Oriented Examples.
package main

import (
	"fmt"
	"strings"
)

type person struct {
	FirstName  string
	LastName   string
	FavFlavors []string
}

type secretAgent struct {
	Person        person
	LicenseToKill bool
}

type stringer interface {
	String()
}

func (p person) String() string {
	var str strings.Builder
	str.WriteString(p.FirstName + "\n")
	str.WriteString(p.LastName + "\n")
	for _, v := range p.FavFlavors {
		str.WriteString(v + "\n")
	}

	return str.String()
}

func (sa secretAgent) String() string {
	var str strings.Builder
	str.WriteString(sa.Person.String())
	if sa.LicenseToKill {
		str.WriteString("Licensed To Kill\n")
	}

	return str.String()
}

func main() {
	p1 := person{
		FirstName: "Clay",
		LastName:  "Dunston",
		FavFlavors: []string{
			"orange sherbert",
			"vanilla",
		},
	}

	sa1 := secretAgent{
		Person: person{
			FirstName:  "James",
			LastName:   "Bond",
			FavFlavors: []string{"rum & coke"},
		},
		LicenseToKill: true,
	}

	fmt.Println(p1)

	fmt.Println(sa1)
}
