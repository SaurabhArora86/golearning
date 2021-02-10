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
	if pop, ok := map1["Haryana"]; ok {
		fmt.Println(pop)
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
}
