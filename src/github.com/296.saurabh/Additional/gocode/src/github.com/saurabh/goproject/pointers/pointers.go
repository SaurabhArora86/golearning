package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	a := 5
	b := &a
	// it is same as var b *int = &a

	// c is Value at address of a
	// i.e * - value at &a - address of a
	// * is called defrencing operator if it is like *<pointer>
	c := *b

	fmt.Println(a, b, c, *&a)
	*b = 15
	fmt.Println(a, b, c, *&a)
	// address represents *int which is different than int
	fmt.Printf("type for var %T and type for address %T\n", a, b)

	//Change val with pointer
	// *b is value at address of a
	*b = 10
	fmt.Printf("A value %d\n", a)

	fmt.Println("\n--another example\n")

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	fmt.Println("\n--Important---")
	//Below works for slices but not for arrays oe structs
	//Since slices points to first object of array
	// If u copy an array to another, it is a copy not a reference passed on like slice
	// map also has same behavior as slice
	p := []int{1, 2, 3}
	q := p
	fmt.Println(p, q)
	p[1] = 9
	fmt.Println(p, q)

}
