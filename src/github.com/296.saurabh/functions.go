package main

import (
	"fmt"
)

//Main is the entrypoint and takes no params and returns no values
func main() {

	fmt.Println("Functions")
	sayHello("Hello", "Saurabh")

	//Another
	gting := "Helya"
	n := "Ramu Kaka"
	greet(&gting, &n)

	//variatic params
	sumval := sum("The sum is", 1, 2, 3, 4)
	fmt.Println(*sumval)

	//Handling more than 1 return type
	d, err := divide(4.0, 5.0)
	if err != nil {
		fmt.Println(err)
		//Exit out of main with return, i.e exit out of application
		return
	}
	fmt.Println("Division is: ", d)

	//Anonymous function
	func(msg string) {
		fmt.Println(msg)
	}("Hello!!!!!")

	//Function passed as a variable
	f := func() {
		fmt.Println("Hello Anonymous f")
	}
	f()
	//IMPORTANT METHOD INVOCATION
	mt := greeter{
		g: "Hello Ji",
		n: "Namaste",
	}
	mt.grw()

	an := anonym()
	an("Hello I was inside anonymous")
}

func sayHello(msg, name string) {
	fmt.Println(msg, name)
}
func greet(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*greeting = "Ok"
	fmt.Println(*greeting, *name)
}

//Here values ...int is taking in variatic params
//values becomes slice
func sum(msg string, values ...int) *int {
	fmt.Println("----> In variatic function")
	fmt.Println(values)
	result := 0
	// _ takes in key and v takes in value
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
	return &result

}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("Cannot divide by Zero")
	}
	return a / b, nil
}

type greeter struct {
	g string
	n string
}

func (mtw greeter) grw() {
	fmt.Println(mtw.g, mtw.n)
}

//Function returning anonymous function
//Below returns an anonymous function :   func(string) is return type of anonym
func anonym() func(string) {
	return func(msg string) {
		fmt.Print(msg)
	}
}
