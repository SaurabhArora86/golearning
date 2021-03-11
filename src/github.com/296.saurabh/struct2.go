package main

import "fmt"

//Point is a struct
type Point struct {
	x int
	y int
}

//Circle as Embedded structs
type Circle struct {
	radius float32
	center *Point
}

func changeMe(pt *Point) {
	//Note here, we dont need (*pt.x) for struct
	//This is syntactic sugar where we dotn need to derefence
	//COmpiler automatically does it
	pt.x = 5
}

func changeIt(pt Point) {
	pt.x = 5
}

func main() {
	p1 := Point{y: 9}
	changeIt(p1)
	//Note here, value of x is still 0
	//Means copy is passed, not reference (unlike maps & slices)
	fmt.Println(p1)
	p2 := &Point{y: 3}
	changeMe(p2)
	fmt.Println(*p2)
	fmt.Println("--------> Circle")
	cr := Circle{4.5, p2}
	fmt.Println(*cr.center)
	fmt.Println(cr.center)
	fmt.Println(cr.center.x)
	fmt.Println(cr)
}
