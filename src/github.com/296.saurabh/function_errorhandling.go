package main

import "fmt"

func main() {
	fmt.Println("Error handling in functions")
	a, b := 5.0, 0.0
	result, err := divide(a, b)
	if err != nil {
		fmt.Println("GOT AN ERROR-->", err, "DEN IS", result)
		return
	}
	fmt.Println(result)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return b, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}
