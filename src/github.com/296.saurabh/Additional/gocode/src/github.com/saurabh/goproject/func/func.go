package main

import "fmt"

func grettings(name string) string {
	return "Hello " + name
}

func sum(num1, num2 int) int {
	return (num1 + num2)
}

func sub(num3 int, num2 int) {
	fmt.Println(num3 - num2)
}

func main() {
	fmt.Println("Hello World")
	great := "GREAT"
	// s := grettings("Saurabh")
	sub(5, 3)

	fmt.Println(grettings("Saurabh"))
	fmt.Println(sum(3, 5))
	greeting := "Hello"
	sirname := "Arora"
	sayGreeting(&greeting, &sirname)
	fmt.Println(sirname)
	s1Greeting(&great)
	//variatic function
	summing("Sum is", 1, 2, 3, 4, 5)
	//Anonymous function func () {} () - last brackets is for immediate calling
	// func() {
	// 	fmt.Println("Hello Anonymous")
	// }()

	//function definition aand variable assignment
	variable := func() {
		fmt.Println("Hey There, I am a var assignment")
	}
	variable()

	// struct and func defined below
	// Function can be used in a context (struct here)
	// these are then called methods
	loc := mystruct{
		location:  "Faridabad",
		shortname: "Fbd",
	}
	loc.locfunc()

}

func sayGreeting(greeting, sirname *string) {
	fmt.Println(*greeting, *sirname)
	*sirname = "Ted"
	fmt.Println("Inside function - " + *sirname)
}

func s1Greeting(great *string) string {
	fmt.Println(*great)
	if *great != "follow" {
		return "Hey"
	}
	return ("Hello guys")
}

//variable length parameter i.e variatic parameter
func summing(msg string, val ...int) {
	fmt.Println(val)
	result := 0
	for _, v := range val {
		result += v
	}
	fmt.Println(msg, result)
}

type mystruct struct {
	location  string
	shortname string
}

// below fucntion is a value receiver since it is receiving values for struct
func (h mystruct) locfunc() {
	fmt.Println(h.location, h.shortname)
}
