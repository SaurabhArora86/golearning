package main

import (
	"fmt"
)

func main() {

	// Typed constant
	const myConst int = 42
	const a float32 = 42.2
	const b string = "Hello Constant"

	//THis wont work
	// myConst = 52 since this is a constant
	fmt.Printf("%v, %T\n", myConst, myConst)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	//Another option is no need to specify type explicitely
	//compiler would automatically assign it
	//this is untyped constant
	const o = 10
	fmt.Println("Hello", o)
	fmt.Printf("%v, %T\n", o, o)

	// COnstants behave like variables only
	const x int = 42
	var y int = 32
	fmt.Printf("%v, %T\n", x+y, x+y)

	//Below wont work
	//var z float32 = 72
	//fmt.Printf("%v, %T\n", z+y, z+y)
	//fmt.Printf("%v, %T\n", x+z, x+z)
	//Error:
	// 	src\github.com\296.saurabh\constants.go:25:26: invalid operation: z + y (mismatched types float32 and int)
	// src\github.com\296.saurabh\constants.go:26:26: invalid operation: x + z (mismatched types int and float32)
}
