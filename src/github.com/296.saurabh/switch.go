package main

import (
	"fmt"
)

func main() {
	a := 20
	//switching on tag
	switch a {
	//You can add multiple validation with same case
	case 10, -1:
		fmt.Println("Less than 10")
	case 20:
		fmt.Println("Equal to 20")
	//Optional
	default:
		fmt.Println("Between 10 and 20")
	}
	//tagless switch
	switch {
	case (a <= 20):
		fmt.Println("Less than equal to 20")
		fmt.Println("Goof work")
	default:
		fmt.Println("Greater than 20")
	}
}
