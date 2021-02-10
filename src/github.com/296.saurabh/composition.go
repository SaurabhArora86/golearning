package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	//Embedded Animal struct to allow composition
	Animal
	Speed  int
	CanFly bool
}

func main() {
	b := Bird{}
	b.Name = "Peacock"
	b.Origin = "India"
	b.Speed = 10
	b.CanFly = false

	fmt.Println(b)
	fmt.Println(b.Name)

}
