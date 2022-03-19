package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
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

}
