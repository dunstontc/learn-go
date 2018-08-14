package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

// another shape
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// implements the shape interface
func (s square) area() float64 {
	return s.side * s.side
}

// implements the shape interface
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Prints the area of a `shape`
func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}
