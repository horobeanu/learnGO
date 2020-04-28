package main

import (
	"fmt"
	"math"
)

// define "geometry" interface
type geometry interface {
	Area() float64
	Perim() float64
}
// now "geometry" can be used as an interface
// and all types that have Area() and Perim() methods
// are automatically implementing "geometry" interface

// To implement an interface in Go,
// we just need to implement all the methods in the interface.
type rectangle struct {
	width, height float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}
func (r rectangle) Perim() float64 {
	return 2*r.width + 2*r.height
}
// now rectangle implements "geometry" interface

// same for circle
type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) Perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type,
// then we can call methods that are in the named interface.
// Here’s a generic measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
