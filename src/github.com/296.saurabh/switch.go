package main

import (
	"fmt"
)

func main() {
	a := 20
	//switching on tag
	switch a {
	//You can add multiple validation with same case
	// switches in go have implicit break
	case 10, -1:
		fmt.Println("Less than 10")
	case 20:
		fmt.Println("Equal to 20")
	//Optional
	default:
		fmt.Println("Between 10 and 20")
	}
	//tagless switch, allows case to have comparison operator
	switch {
	case (a <= 20):
		fmt.Println("Less than equal to 20")
		fmt.Println("Goof work")
	default:
		fmt.Println("Greater than 20")
	}
	//Type case in switch
	//interface{} can take any data type
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("I am integer")
	case string:
		fmt.Println("I am string")
	default:
		fmt.Print("Nothing")
	}
}
