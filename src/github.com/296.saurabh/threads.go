package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello"
	//Add syncronizes the go routine which is an anonymous function below
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg, "anonymous")
		//Done reduces the wait group by 1 so it becomes 0 and then go knows that all routines are finished
		wg.Done()
	}(msg)
	msg = "Message"
	fmt.Println(msg)
	wg.Wait()
}
