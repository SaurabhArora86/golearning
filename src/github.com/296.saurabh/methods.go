package main

import "fmt"

// With structs, copy of objects gets created while passing values to functions
// Unlike maps and slices
//To make it work, pointer ref is passed

type greeting struct {
	greet string
	name  string
}

// func (gp greeting) gr() {
// 	fmt.Println(gp.greet, gp.name)
// }

func (gp *greeting) gr() {
	//No need to specify * while gp.greet since this is handled by compiler itself
	//Now since pointer ref is passed, if we change the value in function,
	//it will be changed outisde as well
	fmt.Println(gp.greet, gp.name)
	gp.name = ""
}

type Student struct {
	name     string
	grades   []int
	age      int
	avgGrade float32
}

func (s Student) getAge() int {
	return s.age
}

func (s Student) setAge(a int) int {
	s.age = a
	return s.age
}

func (s *Student) getAvg() float32 {
	//return s.grades / len(s.grades)
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	s.avgGrade = float32((sum) / len(s.grades))
	return s.avgGrade
}

func main() {
	g := greeting{
		greet: "Hello Ji",
		name:  "Namaste",
	}
	g.gr()
	fmt.Println("Name is: ", g.name)

	s1 := Student{"Saurabh", []int{10, 20, 30}, 99, 0}
	fmt.Println("Age is", s1.getAge())
	//This wont work outside the function, to make it work, pass the pointer
	// as with gr() method
	fmt.Println("Age set is", s1.setAge(10))
	fmt.Println(s1.getAge())
	fmt.Println("Avg Grades:", s1.getAvg())
}
