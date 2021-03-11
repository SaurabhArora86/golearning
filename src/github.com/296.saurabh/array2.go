package main

import "fmt"

func main() {
	//Array declaration
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2}
	//THis wont work --> c := [3]int
	var c [3]int
	fmt.Println("In Arrays")
	fmt.Println("A is", a)
	fmt.Println("B is", b)
	fmt.Println("C is", c)
	c[2] = 4
	fmt.Println("C is", c)
	//fmt.Println("Address:", &a, &b)
	fmt.Print("Arrays are mutable")
	c[2] = 9
	fmt.Println("C is", c)
	//Arrays when assigned to another array, a copy gets created
	// so you can change the new array and wotn imapct previous
	z := a
	z[1] = 5
	fmt.Println("Copied Array", z)
	fmt.Println("orig Array", a)
	//To avoid this
	q := &z
	q[1] = 15
	fmt.Println("Copied Array", *q)
	fmt.Println("orig Array", z)
	//Append cannot work with arrays
	// z = append(z, 99)
	// fmt.Println("orig Array", z)

}
