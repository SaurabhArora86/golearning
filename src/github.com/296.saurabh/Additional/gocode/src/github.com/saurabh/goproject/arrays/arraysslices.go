package main

import "fmt"

func main() {

	// Arrays are of fixed type with return type
	// If u want array of variable length use slices
	var fruitArray [3]string
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"
	fruitArray[2] = "Banana"

	// Declare and Assign at same time
	vegArr := [2]string{"Onion", "Caps"}

	// Slices
	animArr := []string{"Sheep", "Hen", "Cow"}
	fmt.Println("Hello World")
	fmt.Println(fruitArray)
	fmt.Println(fruitArray[1])
	fmt.Println(vegArr)
	fmt.Println(animArr)
	fmt.Println(len(animArr))
	//Slicing like Python
	fmt.Println(fruitArray[1:2])
	fmt.Println(animArr[1:2])
}
