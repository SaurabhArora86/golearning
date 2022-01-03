package main

import (
	"fmt"
	"strconv"
)

// Structs are like classes since there are no classes in go
// You can define a struct and then assign properties liek variables
// and functions

// Defining a struct
// NOTE: Methods are not defined in struct
// NOTE: All values/ keys of Arrays, slice, maps [keys and values] need to be of same data type but not with struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// age       int
	// example []string -- this is a slice
	firstName, lastName, city string
	age                       int
	// to fill values
	//person := Person{firstName: :saurabh .....}
}

//Anonymous struct
// person := struct {firstname string}
// Defined and Initialized at same time
// person := struct {firstname string} {name: "Saurabh"}

//Functions types:Value recievers and pointer receiver
// Value receiver are for methods that doesnt change anything like sum or caluclation etc
// Pointer receivers are fucntions that change something

//greeting method (value reciever)
//Normal func: func funcname (parameters) returntype {}
// func (identifier <struct>) funcname returntype {}
func (p Person) greet() string {
	return ("Hello " + p.firstName + " " + p.lastName + " " + strconv.Itoa(p.age))
}

// changeMe is a pointer receiver
// func (identifier *<struct>) funcname returntype {}
func (p *Person) changeMe() {
	p.age++
}

func main() {

	fmt.Println("In Struct")
	// initialize person usign struct
	person1 := Person{firstName: "Saurabh", lastName: "Arora", city: "Faridabad", age: 26}
	//Another definition
	person2 := Person{firstName: "Chancal", lastName: "Mann", city: "Faridabad", age: 26}
	//Another way to fill
	person3 := Person{"Vicky", "Arora", "Faridabad", 34}
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	// Use . syntax to retrieve value of a field
	fmt.Println(person2.lastName)
	//Changes/ Overrides are allowed
	person2.firstName = "Henry"
	fmt.Println(person2.firstName)
	fmt.Println(person2)
	// using value receiver
	person1.changeMe()
	fmt.Println(person1.greet())

	// NOTE: for maps and slices, pointer reference is passed when one map is passed to another
	// so if we change map in one reference, it would change in another as well
	// where as for Struct, only a copy is passed hence if u make changes in one struct, it wont impact other for eg:
	type Amph struct {
		firstname, lastname string
	}
	male1 := Amph{firstname: "Sarah", lastname: "HEllo"}
	male2 := male1
	male2.firstname = "Hulu"
	fmt.Println(male1)
	fmt.Println(male2)

}
