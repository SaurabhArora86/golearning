package main

import "fmt"

//NOTE : There is no while loop

func main() {
	fmt.Println("Loops")
	//First loop
	// i is limited to scope of loop as usual
	for i := 5; i < 7; i++ {
		fmt.Println("Value of i:", i)
	}
	p := 10
	for ; p < 12; p++ {
		fmt.Println(p)
	}
	q := 12
	//below works both ways
	//for ;a < 14;
	//	{fmt.Println(q)
	//	q++} OR below
	for q < 14 {
		fmt.Println("Value of q", q)
		q++
	}
	// as like other languages
	//continue will take control back to begining of for loop
	// and will not print fmt.Println(i)
	for i := 2; i < 8; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("i is odd", i)
	}
	//collection with loops
	s := [4]int{1, 2, 3, 4}
	for _, p := range s {
		fmt.Println("Value of slice:", p)
	}
	// In case u only want keys
	// it is special
	o := map[string]int{"Name": 1, "age": 4}
	for k := range o {
		fmt.Println(k)
	}
	l := "Hello!!"
	for k, v := range l {
		fmt.Println(k, string(v))
	}

}
