package main

import (
	"fmt"
)

//Pointer that is not initialized holds the value nil for eg:
// var a *int, here value of a is nil and a is a pointer to an integer type
//Unlike Arrays and struct both maps and slices are internally referencing pointer to the original value
//when you copy data to another map/slice
type myStruct struct {
	foo int
}

func main() {

	var a int = 50
	var b *int = &a    // OR b := &a b is a pointer of type int
	fmt.Println(a, *b) // * is defrencing operator
	fmt.Println(&a, b)
	a = 27 // *b = 27
	fmt.Println(a, *b)

	fmt.Println("Another example")
	p := []int{1, 2, 3}
	x := &p[0]
	y := &p[1]
	fmt.Printf("%v %p %p\n", p, x, y)

	//Struct example
	// var ms *myStruct
	ms := &myStruct{foo: 42}
	//Below tells ms is holding address of object that has value 42 in it
	fmt.Println(ms)
	fmt.Println(*ms)

	fmt.Println("Another way for struct, by default it holds address of struct")
	ms1 := new(myStruct)
	fmt.Println(ms1)

	fmt.Println("Another one----")
	var rr *myStruct
	//Pointer of type myStruct holds nil as value
	fmt.Println(rr)

	//Another way of initializing a struct without passing any params for it
	//initializes struct with  default values of parms datatype (for params in struct)
	fmt.Println("Another one----using new")
	//This is essentially same as ms := &myStruct{}
	pk := new(myStruct) //new can only be used with blank value
	fmt.Println(pk)
	fmt.Println(*pk)
	(*pk).foo = 55 //() is requried since * has lower preference than .
	// OR pk.foo also works (this is syntatic sugar)
	fmt.Println(pk.foo)

	pk1 := &myStruct{}
	fmt.Println(pk1)

	//maps as example
	fmt.Println("Maps as example")
	m1 := map[int]int{1: 20, 2: 30}
	m2 := m1
	fmt.Println(m1)
	m1[1] = 40
	fmt.Println(m1)
	fmt.Println(m2)
}
