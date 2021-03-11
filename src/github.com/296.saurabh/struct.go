package main

import "fmt"

//Unlike maps and slices, when you copy a struct, it actually copies not reference

type Doctor struct {
	//NOTE: YOu need to capitalize first letter to make fields visible outside package
	// like Num not num
	num        int
	actorName  string
	companions []string
}

//Another one
type mystruct struct {
	name string
	jobs []string
}

func main() {
	doc := Doctor{
		num:       5,
		actorName: "Saurabh",
		companions: []string{
			"Lisa",
			"Ray",
		},
	}
	fmt.Println(doc)
	fmt.Println(doc.actorName)
	fmt.Println(doc.companions[1])
	fmt.Println("------->")
	//YOu can set value for particular fields out af all in strcut and
	// remaining becomes defaults for eg num becomes 0
	docf := Doctor{
		num:       5,
		actorName: "Pikachu",
	}

	fmt.Println(docf)

	fmt.Println("Another way")
	doc2 := Doctor{
		5,
		"Saurabh",
		[]string{
			"Lisa",
			"Ray",
		},
	}
	fmt.Println(doc2)
	fmt.Println(doc2.num)
	fmt.Println(doc2.companions[0])
	fmt.Println("Possible to ignore a particular value from struct")

	doc3 := Doctor{
		num:       5,
		actorName: "Saurabh",
	}
	fmt.Println(doc3)
	fmt.Println("Anonymous struct")
	pdoc := struct {
		name string
		age  int
	}{
		name: "Saurabh", age: 24,
	}
	fmt.Println(pdoc)
	fmt.Println("Copies are passed")
	gdoc := pdoc
	gdoc.age = 29
	fmt.Println(gdoc)
	fmt.Println(pdoc)
	fmt.Println("To change original struct")
	kdoc := &gdoc
	kdoc.age = 33
	fmt.Println(*kdoc)
	fmt.Println(gdoc)

	mystrval := mystruct{
		name: "Saurabh",
		jobs: []string{
			"Homealone",
			"Multiple",
		},
	}

	fmt.Println(mystrval)
	fmt.Println(mystrval.jobs)

}
