package main

import "fmt"

//IMP: slices can be defined with make
//a := make([]int,len,cap)
//a :=make([]string,3) <--cap is optional

func main() {
	fmt.Println("Slices")
	a := []int{1, 2, 3}
	var b []int
	fmt.Println("A is", a)
	fmt.Println("B is", b, "Length is", len(b))
	// IMP: Slices unlike arrays refernce to same underlying data
	fmt.Println("--Referencing")
	s := a
	s[1] = 25
	fmt.Println("A is", a)
	fmt.Println("s is", s)
	//Another way to create slices
	//Below works with arrays as well
	y := []int{1, 2, 3, 4, 5, 6, 7, 8}
	u := y[3:]
	o := y[2:6]
	fmt.Println("U is", u)
	fmt.Println("O is", o)
	fmt.Println("Usign MAKE")
	l := make([]int, 3)
	fmt.Println(l)
	fmt.Println("Addition to slice using append")
	//Append doesnt work with arrays since they are of fixed size
	// append(slice,one or more elements to append)
	y = append(y, 89, 90, 100)
	fmt.Println(y)
	// Removing element is best usign loops
}
