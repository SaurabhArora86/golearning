package main

import (
	"fmt"
	"math"
)

// INterfaces are meant only for methods, not variable types
type shape interface {
	area() float32
}

type boundary interface {
	perimeter() float32
}

// Circle and Rectangle internally inherits shape
// Since area method is implemented by both circle and rectangle
type circle struct {
	radius float32
}

type rectangle struct {
	length  float32
	breadth float32
}

// Method name, params and return tyoe must match as declared in interface
func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

//Example, one without pointer and other with pointer
func (r *rectangle) area() float32 {
	return r.length * r.breadth
}

//Normally to implement iterface on ground, you need helper functions that hides
// the underlying implementation
func measure(sh shape) {
	fmt.Println(sh.area())
}

func (r *rectangle) perimeter() float32 {
	return 2 * (r.length * r.breadth)
}
func main() {
	c := circle{3}
	r := rectangle{3, 4}
	fmt.Println("Measuring shape")
	s := []shape{c, &r}
	// IMP: since the common thing between circle and rectangle is area, with s as a slice
	//of type shape, you can only access area using s, since shape interface only defines area
	measure(c)
	//For traversing in slice
	for _, v := range s {
		fmt.Println("Printing area with range")
		fmt.Println(v.area())
	}
	p := []boundary{&r}
	for _, v := range p {
		fmt.Println(v.perimeter())
	}
}
