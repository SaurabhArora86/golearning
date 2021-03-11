package main

import "fmt"

// Interfaces are used on structs or other datatype when they have common behavior
// and you want to impose that behavior
// i.e they have to implement that common function
// NOTE interfaces are meant for methods only

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length, breadth float64
}

//since circle and rectangle are of type shape
//area() needs to be implemented by both circle and rectangle
// since it is a function declared in interface in
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

//helper function
func getArea(sh shape) float64 {
	return (sh.area())
}

func main() {
	fmt.Println("Interfaces")
	c := circle{3.0}
	r := rectangle{3.0, 4.0}
	i := []shape{c, r}
	for _, v := range i {
		fmt.Println(v.area())
		fmt.Println(getArea(v))
	}
}
