package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Hello Channels")
	// Channel Creation - strongly typed
	//for Buffer channel use ch :=make(chan int, num)
	ch := make(chan int)
	wg.Add(6)
	go func() {
		i := <-ch
		fmt.Println("Hello Number: ", i)
		wg.Done()
	}()

	go func() {
		ch <- 42
		wg.Done()
	}()

	//ANother example where each function is either a receiver or sender
	// <-chan corresponds to sender and chan<- is channel data receiver
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println("Second example", i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 33
		wg.Done()
	}(ch)

	fmt.Println("Another example")
	go func(ch <-chan int) {
		//First alternative
		// for i := range ch {
		// 	fmt.Println("Third example", i)
		// }
		//Second alternative
		for {
			if i, ok := <-ch; ok {
				fmt.Println("Third example", i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 99
		ch <- 100
		//It is imp to close channel otherwise the upper go function will get into deadlock
		//Since it would keep on waiting for more data
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
