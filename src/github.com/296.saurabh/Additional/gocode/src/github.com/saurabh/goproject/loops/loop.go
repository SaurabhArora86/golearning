package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//init and post statements are optional
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}
