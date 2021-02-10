package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func main() {
	//Below gives number of cores processor has
	fmt.Println("Threads %v\n", runtime.GOMAXPROCS(-1))
	for i := 0; i < 3; i++ {
		wg.Add(2)
		//Read lock
		m.RLock()
		go hello()
		//Normal lock so that noone can access
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func hello() {
	fmt.Printf("In Hello %v\n", counter)
	//Releaseing read lock
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
