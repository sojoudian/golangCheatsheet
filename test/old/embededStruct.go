package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type bestProgrammingLanguage struct {
	person
	itb bool
}

func main() {

	b1 := bestProgrammingLanguage{
		person: person{
			first: "Ken",
			last:  "Thompson",
			age:   78,
		},
		itb: true,
	}

	p1 := person{
		first: "Richard",
		last:  "stallman",
		age:   68,
	}
	p2 := person{
		first: "Dennis",
		last:  "Ritchie",
		age:   70,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	fmt.Println(b1.first, b1.last, b1.age, b1.itb)

}
