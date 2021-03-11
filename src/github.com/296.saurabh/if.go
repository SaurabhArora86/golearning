package main

import (
	"fmt"
)

func main() {

	if true {
		fmt.Println("Hello There")
	}

	map1 := map[string]int{
		"Delhi":   10,
		"Haryana": 20,
	}
	//Initializer syntax
	//Pop scope is limited to if
	if pop, ok := map1["Haryana"]; ok {
		fmt.Println(pop)
	}
	if pop1, ok := map1["Jhujjar"]; ok {
		fmt.Println(pop1)
	}

	if _, ok := map1["Pun"]; ok {
		fmt.Println("Missing")
	}

	number := 20
	if number < 20 {
		fmt.Println("Less then 20")
	} else if number == 20 {
		fmt.Println("Number equal")
	} else {
		fmt.Println("Number greater than 20")
	}

	no := 30
	if no < 20 {
		fmt.Println("Lesst than 20")
	} else if no > 20 {
		fmt.Println("Greater thn 20")
	} else {
		fmt.Println("Equal to 20")
	}
}
