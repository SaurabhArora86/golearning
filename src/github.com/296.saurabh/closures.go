package main

import "fmt"

//Closure are inner functions that holds the value of a variable just outside its scope
func anonym() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	returnedI := anonym()
	//Note below, returnedI keeps a track of I
	fmt.Println(returnedI())
	fmt.Println(returnedI())
	fmt.Println(returnedI())
	fmt.Println("Hello")
}
