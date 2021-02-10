package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello Interfaces")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 3; i++ {
		fmt.Println((inc.Increment()))
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

// concrete object
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//Another example

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (in *IntCounter) Increment() int {
	*in++
	return int(*in)
}
