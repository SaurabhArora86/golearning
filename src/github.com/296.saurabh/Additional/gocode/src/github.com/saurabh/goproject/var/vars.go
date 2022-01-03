package main

import "fmt"

var name = "saurabh"
var age = 10
var isCool = true
var floating = 2.3

const myName = "Vicky"

func main() {
	//Short hand for variables (only works inside function)
	sirname := "Arora"
	city, userid := "Faridabad", "296.saurabh"
	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", age)
	fmt.Println(isCool)
	fmt.Println(myName)
	fmt.Println(sirname)
	fmt.Printf("%T\n", floating)
	fmt.Println(city, userid)
	fmt.Println("Printing Pichu and Kichu")
	var pichu string
	fmt.Println(pichu)

	var kichu int
	fmt.Println(kichu)
}
