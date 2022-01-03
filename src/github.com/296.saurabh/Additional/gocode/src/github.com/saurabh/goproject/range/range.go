package main

import "fmt"

func main() {
	//Range works with Maps, slices and arrays
	//slice
	id := []int{3, 4, 5, 6, 7, 10}

	for i, val := range id {
		fmt.Printf("Index is %d and Value is %d\n", i, val)
	}
	fmt.Println("Second important")
	// If u dont want to use index use _
	for _, val := range id {
		fmt.Printf("Value is %d\n", val)
	}
	fmt.Println("Third important")
	// If u dont define index at all
	// val gets index value
	for val := range id {
		fmt.Printf("Value is %d\n", val)
	}
	for sum, val := range id {
		sum += val
		fmt.Printf("Sum: %d\n", sum)
	}

	//Using maps
	email := map[string]string{"Saurabh": "saurabh@gamil.com", "vicky": "saurabh@gamil.com"}

	for k, v := range email {
		fmt.Printf("%s: %s\n", k, v)
	}
}
