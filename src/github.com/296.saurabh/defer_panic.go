package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("First")
	//defer defers the execution till the last statement of function is executed
	// then runs defer jsut before returning back
	defer fmt.Println("Middle")
	fmt.Println("End")
	//Defer works in reverse order
	defer fmt.Println("Last")

	fmt.Println("Another one")
	// Here defer will print start coz it considers the argument at the time of call not at the time of execution
	a := "start"
	defer fmt.Println(a)
	a = "end"

	//Panic by go runtime
	// x, y := 1, 0
	// fmt.Println(x / y)

	//Panic purposeful

	// fmt.Println("Start execution")
	// panic("Something happened")
	// fmt.Println("End execution")

	//Panic again
	//IMP NOTE: Panic will execute after defer
	panic("Problem happened")
	_, err := os.Create("/tmp/file")
	if err != nil {
		//log.Fatal(err)
		panic(err)

	}
}
