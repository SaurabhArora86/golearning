package main

import "fmt"

// GO doesnt support Inheritance but it can be achieved using composition
// i.e HAS A relationship but not a IS-A relationship of inheritance
// By embedding one struct into another (Struct is like a class in JAVA)

type Animal struct {
	name string
}

type Cap string

// Animal enbedded into Mammal
type Mammal struct {
	Animal
	canFly bool
	speed  int
}

func main() {
	fmt.Println("Hello Compostion")
	c := Mammal{}
	c.name = "Cat"
	c.canFly = false
	c.speed = 20
	fmt.Println(c)

	fmt.Println(c.name)
	//However below doesnt work (sicne not a IS-A relationship)
	//b := Mammal{name: "Penguin", canFly: false, speed: 10}
	// To make it work use:
	k := Mammal{
		Animal: Animal{name: "Tiger"},
		canFly: false,
		speed:  40,
	}
	fmt.Println("\n --Compostions vs Inheritance")
	fmt.Println(k)

}
