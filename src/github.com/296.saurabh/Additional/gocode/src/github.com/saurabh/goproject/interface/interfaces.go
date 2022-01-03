package main

import (
	"fmt"
	"math"
)

// Define interface

type Shape interface {
	area() float64
}

// Define Circle and Rectangle Struct

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// Circle implements interface Shape
// Implicit implementation no need of using any extend or anything
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Get area is a common function. ANyone that implements interface (Interface Type) can invoke getArea function
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))

	//Another example implementation(details below)
	// This is how interfaces are used; var <variable> interfaceName = <type> -- type is what implements interface
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello GO!"))

	// Third example
	var i I = T{"Hello"}
	i.M()

	// Type Assetions --------------->
	// Defining an interface with value as hello
	//t := i.(T) statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t
	var y interface{} = "Hello"

	p := y.(string)
	fmt.Println(p)

	q, ok := y.(float64)
	fmt.Println(q, ok)

	// Type Assertiosn another ------------------>
	do(21)
	do("hello")
	do(true)

}

//--- Another Example----

type Writer interface {
	//This is an interface
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
	// This is blank struct
}

// ConsoleWriter implements Writer interface and defines its method signature
func (cn ConsoleWriter) Write(data []byte) (int, error) {
	r, error := fmt.Println(string(data))
	return r, error
}

//--- Another Example----

type I interface {
	M()
}

type T struct {
	S string
}

// M - This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

// Type Switches
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
