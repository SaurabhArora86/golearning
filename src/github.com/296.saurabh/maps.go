package main

import (
	"fmt"
)

//Key Value pairs
// Syntax mapname := map[keydatatype]valuedatatype{keys: values}
//eg : map1 :=map[string]string{"Hello": "Saurabh",}
// Key cannot be slices, maps and functions since thye have to be able to be tested for equality
// Keys need to be unique, if u try to add again, value will be updated
// Return order of map is not guaranteed whereas arrays and slices, order is guaranteed
// Just like slices, if u copy a map to another map and make changes in other map, original will also change
// i.e reference is passed, it doesnt create a copy of map
func main() {

	population := map[string]int{
		"delhi":   100,
		"Haryana": 200,
	}
	fmt.Println(population)
	fmt.Println("Deleting map values")
	delete(population, "delhi")
	fmt.Println(population)

	statepopulation := make(map[string]int)
	fmt.Println(statepopulation)
	fmt.Println("Filling values")
	statepopulation = map[string]int{
		"cal": 8,
	}
	fmt.Println(statepopulation)
	fmt.Println(statepopulation["cal"])
	fmt.Println("Adding values")
	statepopulation["ohio"] = 10
	fmt.Println("Updating values")
	statepopulation["cal"] = 9
	fmt.Println(statepopulation)

	//Another example
	students := make(map[string]int)
	students["b"] = 5
	students["c"] = 10
	students["a"] = 3
	fmt.Println(students)
	delete(students, "a")
	fmt.Println(students["b"])
	fmt.Println("Checking key exits")
	_, ok := students["ohio"]
	fmt.Println(ok)
	_, ok = students["c"]
	fmt.Println(ok)
	//Deleting in copied mapwill delete delete values in original, reference is passed
	st := students
	delete(st, "c")
	fmt.Println(students)
	st["d"] = 5
	fmt.Println(students)

}
