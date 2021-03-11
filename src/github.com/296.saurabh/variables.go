package main

import (
	"fmt"
	"strconv"
)

//declaration at package level cannot have := syntax
// it has to eb declared full
var p float32 = 9.9

//"If u declare a variable with capital letter than it is available outside the package
// otherwise with lower case, scope of variables is limited to all files within package"
// for package level, it has to be declared at package level---here"
// if it is main() then it is blacoked scope and not available outside block
var Q string = "Hello Packages"

// var can be clubbed for multiple variable declaration
var (
	counter int     = 1
	ratio   float32 = 32
	name    string  = "Ramu"
)

func main() {
	var i int = 50 //first form
	var j int      //second form in case assignment will happen in a loop or somewhere else
	j = 20
	f := 52 //third one, compiler internally identifies the type of variable
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(j)
	j = 30 //can change value
	fmt.Println(j)
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", p, p)
	fmt.Println(name)

	//Note tha variable declared inside main takes precedence over global variable
	// This is called shadowing
	fmt.Println("---Shadowing")
	fmt.Println(p)
	var p float32 = 10.0
	fmt.Println(p)

	//Acronymns as a best practise in golang, should be upper case
	//Note camel casing as well
	var theHTTPRequest string = "https://google.com"
	fmt.Println(theHTTPRequest)

	//shadowing concept : variable at inner most scope takes precedence
	var name string = "Kaka" //declaration again
	counter = 100            //reassignment
	fmt.Println(name)
	fmt.Println(counter)
	fmt.Println(Q)
	//conversion
	fmt.Println("--conversion")
	counter2 := float32(counter)
	fmt.Printf("%v,%T", counter2, counter2)
	fmt.Println("====String conversion IntergerToAscii")
	counter3 := strconv.Itoa(counter)
	fmt.Printf("%v,%T", counter3, counter3)
}
