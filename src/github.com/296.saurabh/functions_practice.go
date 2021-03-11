package main

import "fmt"

func main() {
	fmt.Println("Inside functions")
	greeting := "Hello"
	name := "Stacy"
	greet(greeting, name)
	addressgreet(&greeting, &name)
	fmt.Println(name)

	//Variatic params
	fmt.Println("Variatic params has to be last and are passed as a slice to function")
	val := sum1("Result", 1, 2, 3)
	fmt.Println(val)
	fmt.Println("Return type as pointer")
	val2 := sum2("Result", 1, 2, 3)
	fmt.Println(*val2)
	fmt.Println("Anonymuous function along with functions as type")
	// func() {
	// 	fmt.Println("I am anonymuous")
	// }()
	fn := func() {
		fmt.Println("I am anonymuous")
	}
	fn()

}

func greet(greeting, name string) {
	fmt.Println(greeting, name)
}

func addressgreet(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Tommy"
}
func sum1(msg string, values ...int) int {
	result := 0
	for _, v := range values {
		//fmt.Println("Value is", v)
		result += v
	}
	return result
}
func sum2(msg string, values ...int) *int {
	result := 0
	for _, v := range values {
		//fmt.Println("Value is", v)
		result += v
	}
	return &result
}
