package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//Methods are simply Function reciever
// func (receiver obj) funcName funcReturnType {}
func (v Vertex) abc() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//This is a normal function and not a receiver
func xyz(q Vertex) float64 {
	return math.Sqrt(q.X*q.X + q.Y*q.Y)
}

// NOTE: methods as function receiver can be non struct type too
type Lamb float64

func (L Lamb) sample() float64 {
	return float64(-L)
}

// -- Difference between pointer and value receiver
// Value receiver works on a copy of Vertex hence the actual value is not changed.
// Where as pointer reviever is used to actually change the values
//Change the pointer receiver to value receiver and see the values

type Ptex struct {
	X, Y float64
}

// Value receiver
func (v Ptex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Pointer receiver

func (v *Ptex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Below becomes the value receiver, you can see teh difference with uncommenting below code
// func (v Ptex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

//With a value receiver, the Scale method operates on a copy of the original Vertex value.
// (This is the same behavior as for any other function argument.)
//The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

func main() {
	P := Vertex{3, 4}
	// Function call is made usign object of struct since the fucntion is function receiver
	fmt.Println(P.abc())
	fmt.Println("Normal fucntion call is different")
	fmt.Println(xyz(P))
	K := Lamb(10.2)
	fmt.Println("Method as non struct type")
	fmt.Println(K.sample())
	fmt.Println("Difference between pointer and receiver fucntion")
	v := Ptex{3, 4}
	// GO interprets below as (&v).Scale(10) since the pointer receiver takes in the pointer
	v.Scale(10)
	fmt.Println(v.Abs())
}
