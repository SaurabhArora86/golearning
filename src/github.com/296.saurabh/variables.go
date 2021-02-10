package main

import "fmt"

//declaration at package level cannot have := syntax
// it has to eb declared full
var p float32 = 9.9

//"If u declare a variable with capital letter than it is available outside the package
// otherwise with lower case, scope of variables is limited to all files within package"
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
	f := 52 //compiler internally identifies the type of variable
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(j)
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", p, p)
	fmt.Println(name)

	//Acronymns as a best practise in golang, should be upper case
	//Note camel casign as well
	var theHTTPRequest string = "https://google.com"
	fmt.Println(theURL)

	//shadowing concept : variable at inner most scope takes precedence
	var name string = "Kaka" //declaration again
	counter = 100            //reassignment
	fmt.Println(name)
	fmt.Println(counter)
	fmt.Println(Q)
}
