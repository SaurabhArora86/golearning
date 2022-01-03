package main

import "fmt"

// Closures are anonymous functions
// i.e functions with no names
func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := addr()
	for i := 0; i < 5; i++ {
		fmt.Println(sum(i))
	}
	fmt.Printf("Typeof sum: %T\n", sum)

}
