package main

import "fmt"

func main() {
	x := 10
	y := 10

	// NOTE the use of {} instead of () with if
	if x < y {
		fmt.Printf("%d is smaller than %d\n", x, y)
		fmt.Printf("%T is the type of x\n", x)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", y, x)
	} else {
		fmt.Printf("%d is smaller than %d\n", y, x)
	}
	// If can have a short statement or assigment before the actual if statement
	if p := 5; p < y {
		fmt.Printf("I am inside if %v\n", p)
	}
}
