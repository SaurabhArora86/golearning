package main

import "fmt"

func main() {
	var n bool = true
	p := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", p, p)

	//NOTE: Default value of variables
	var i int
	var m bool
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", m, m)

	a := 20
	b := 10
	c := 5

	fmt.Println(a + b)
	fmt.Println(a % b)
	fmt.Println(a / b)
	fmt.Println(c / b)

	//float
	f := 32.5
	fmt.Printf("%v, %T\n", f, f)
	var g float32 = 32.5
	fmt.Printf("%v, %T\n", g, g)

	//string
	s := "Hello String"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s)
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), s)
	//NOTE Strings are immuntables
	// i.e s[2] = "u" wont work

	//String concat
	s2 := " I am well known"
	fmt.Println(s + s2)
	fmt.Printf("%v, %T\n", s+s2, s+s2)

}
