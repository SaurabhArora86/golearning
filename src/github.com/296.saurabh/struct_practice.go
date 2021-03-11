package main

import "fmt"

//Structs are value types so unlike maps and slices
//copies gets passed and you can change without changing original struct

type mystruct struct {
	name       string
	othernames []string
}

//Another
type yourstruct struct {
	age    int
	months int
	year   int
}

// Using struct
//type name struct {name string,age int}
//Anonymous doc :=struct{name string,age int}{Saurabh,20,}

func main() {
	fmt.Println("In struct prac")
	struct1 := mystruct{
		name: "Saurabh",
		othernames: []string{
			"awsome",
			"welcome",
			"both",
		},
	}
	fmt.Println(struct1)
	fmt.Println(struct1.othernames[1])
	//using positional variables
	struct2 := mystruct{
		"Vicky",
		[]string{
			"Awsome",
			"Pikuchu",
		},
	}
	fmt.Println(struct2)
	yourstruct1 := yourstruct{
		age:    10,
		months: 6,
		year:   0,
	}
	fmt.Println(yourstruct1.age)
}
