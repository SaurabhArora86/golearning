package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

// Bird has animal characteristics - composition
type Bird struct {
	Animal
	Speed int
}

func main() {
	fmt.Println("Embedded structs")
	//Initializing Bird
	b := Bird{}
	b.Name = "Hello"
	b.Origin = "Australia"
	b.Speed = 50
	fmt.Println(b)
	//Another way
	c := Bird{
		Animal: Animal{Name: "Way", Origin: "India"},
		Speed:  50,
	}
	fmt.Println(c)
}
