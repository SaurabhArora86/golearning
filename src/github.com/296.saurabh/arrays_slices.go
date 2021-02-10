package main

import (
	"fmt"
)

func main() {

	//Array declaration
	// Notes:
	// 	Collection of items of same type
	// 	Fixed size
	// 	Declaration style:
	// 		a :=[3]int{1,2,3}
	// 		a :=[...]int{1,2,3}
	// 		var a[3]int //Array of 3 integers each with value 0

	var students [3]string
	fmt.Println(students)
	students[0] = "Kitty"
	students[2] = "Mitty"
	fmt.Println(students)
	fmt.Println(students[1])
	fmt.Println(students[0])
	fmt.Println(len(students))
	//Below doesnt work
	//var classes [...]string

	//declaration and assignment
	//syntax eg: arrayname :=[number]datatype{value1, value2}
	gr := [3]int{1, 12, 123}
	//create an array of variable size
	grade := [...]int{22, 33, 44}
	fmt.Println(gr)
	fmt.Println(grade)
	//fmt.Println(classes)
	//NOTE: Below will throw an error:
	//invalid array index 3 (out of bounds for 3-element array)
	// grade[3] = 55
	// fmt.Println(grade)
	// thats why we use slices

	//NOTE in golang if u assign an array, a copy of array gets created (not the pointer)
	b := students
	b[1] = "Muke"
	fmt.Println(b)
	//TO AVOID ABOVE
	c := &students
	c[2] = "Kite"
	fmt.Println(*c)
	fmt.Println(students)

	//SLICES - BELOW will throw an error since the array sized was fixed
	// at inititiation time
	// grade[3] = 44
	// fmt.Println(grade)
	//Here comes the concept of slices
	//DECLARATION and INITIALIZATION : name :=[]string{1,2,3}
	// Slices are backed by arrays that is go uses arrays underlying
	slice1 := []string{"He", "She", "We"}
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	fmt.Println("Blank slice")
	slice99 := []int{}
	// slice99[1] = 10
	fmt.Println(slice99)
	//IMPORTANT: SLICES are naturally reference unlike arrays
	// i.e if u pass a slice to another var and change a value in new variable
	//original data of slice also changes (unlike arrays where & is required for that
	// otherwise a copy of array gets created)
	slice2 := slice1
	slice2[1] = "Changed"
	fmt.Println(slice2)
	fmt.Println(slice1)

	//Another way to initialize slices
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice4 := slice3[:]   //slice of all elements
	slice5 := slice3[:3]  //slice from first element till 3rd element i.e from inde 0 till index 2
	slice6 := slice3[3:]  //slice from 4th element i.e from index 3 till end
	slice7 := slice3[4:6] //slice of 5th and 6th element i.e from index 4 till index 5
	//first value (index) is inclusive and last value is exclusive
	fmt.Println(slice4)
	fmt.Println(slice5)
	slice4[4] = 33
	fmt.Println(slice6)
	fmt.Println(slice7)
	//Below will throw an error
	// slice3[8] = 10
	// fmt.Println(slice3)
	//WAY TO DO IT is append function
	fmt.Println("Addition to slice")
	slice3 = append(slice3, 10)
	fmt.Println(slice3)
	//Another way to dclare a slice
	// name = make([]type, length, capacity) : Capacity is optional
	l := make([]int, 3, 10)
	fmt.Println(l)
	fmt.Println(len(l))
	fmt.Println(cap(l))
	l[1] = 44
	//Multiaddition to slice at once
	fmt.Println("Note here append adds to end so 4th, 5th, 6th and 7th element")
	l = append(l, 2, 3, 4)
	fmt.Println(l)
	fmt.Println(len(l))
	fmt.Println(cap(l))
	fmt.Println("Concatenation of slices syntax")
	l = append(l, []int{66, 77, 88, 88}...)
	fmt.Println(l)
	fmt.Println("Addition is using append, for popping use loop and create another slice")
	sSlice := append(l, 1, 2)
	fmt.Println("S Slice", sSlice)
	//Copy
	mySlice := make([]int, 3, 5)
	copy(mySlice, l)
	fmt.Println(mySlice)
	fmt.Println("l slice ", l)

}
