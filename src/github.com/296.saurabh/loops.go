package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	// Here j is scoped to the main function not limited to loop
	j := 10
	for ; j < 12; j++ {
		fmt.Println(j)
	}
	//Since Golang doesnt have while or do while
	//Below performs while loop as well
	for j < 13 {
		fmt.Println(j)
		j++
	}
	s := []int{1, 2, 3}

	//Here range gets k as key and v as value
	for k, v := range s {
		fmt.Println(k, v)
	}

	population := map[string]int{
		"delhi":   100,
		"Haryana": 200,
	}
	for k, v := range population {
		fmt.Println(k, v)
		fmt.Println(k)
	}
}
