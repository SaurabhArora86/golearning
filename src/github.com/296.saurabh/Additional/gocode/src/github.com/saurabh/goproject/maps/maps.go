package main

import "fmt"

// Return order of map is not guaranteed

func main() {
	//Define Map
	// make(map[key data type]value data type)
	emails := make(map[string]string)

	emails["Bod"] = "Bod@gmail.com"
	emails["Henry"] = "Henry@gmail.com"
	emails["Mike"] = "Mike@gmail.com"

	//Declare mapand key values
	secemail := map[string]string{"saurabh": "saurabh@gamil.com", "vicky": "saurabh@gamil.com"}
	//Appened goes int map at top
	secemail["kichi"] = "kichi@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bod"])

	//Delete
	// NOTE: If u delete a value that is int, and you try to print that key, it gives 0
	// Below key is printed for string value and it is blank
	delete(emails, "Mike")
	fmt.Println("--deletion--")
	fmt.Println(emails)
	fmt.Println(emails["Mike"])
	fmt.Print("ANother email")
	fmt.Println(secemail)

	//To get only key and values
	for k, v := range secemail {
		fmt.Println(k, v)
	}
	//To get only Keys
	for k := range secemail {
		fmt.Println("Name: " + k)
	}
	/// ---ANother example--------
	// map syntax
	// var = map[keytype]<val type> {
	// key: value,
	// key2: value2, <NOTE , at end of map>
	//	}
	myMap := map[string]string{
		"name":    "saurabh",
		"sirname": "Arora",
		"city":    "faridabad",
	}
	//NOTE: map keys need to pass equiality check
	// Dynamic values in keys are not allowed like slices, maps
	// but arrays etc are allowed
	map2 := map[[2]int]int{}

	//map defined and then filled
	map3 := map[string]int{}
	map3["key"] = 55
	fmt.Println("I am in map main")
	fmt.Println(myMap, map2, map3)
	fmt.Println(map3["key"])
	// IMP: If you mispelled the map key it gives blank or 0 as value
	// to overcome this you can use ok when you are not sure of keys, eg below
	_, ok := map3["ky"]
	fmt.Println(ok)
	fmt.Println(len(map3))
}
