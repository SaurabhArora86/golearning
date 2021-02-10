package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 32
	var j float32 = 42.5

	var p = float32(i)
	var q = int(j)

	var r int
	r = i
	fmt.Printf("%v, %T\n", r, r)
	//above wont work for r = j because of type conversion issue
	//you have to explicitely convert
	r = int(j)

	fmt.Printf("%v, %T\n", p, p)
	fmt.Printf("%v, %T\n", q, q)
	fmt.Printf("%v, %T\n", r, r)

	var str string
	str = strconv.Itoa(i) //Integer to ascii
	fmt.Printf("%v, %T\n", str, str)
	// var con string
	// con = i
	// fmt.Printf("%v, %T\n", con, con)
	//NOTE: Above code will throw error
	//ERROR: cannot use i (type int) as type string in assignment
}
