package main

import "fmt"

func main() {
	fmt.Println("Practise Methods")
	fmt.Println("Methods are functions in known context")
	gt := godtypes{
		name:  "Angel",
		place: []string{"onearth", "heaven"},
		age:   3200,
	}
	gt.callfunc()
	fmt.Println(gt.name)
	fmt.Println(gt.place[1])
}

type godtypes struct {
	name  string
	place []string
	age   int
}

//g is pointer receiver
func (g *godtypes) callfunc() {
	fmt.Println(g.name)
	g.name = "Devil"
	g.place[1] = "Devil in Devil"
}
