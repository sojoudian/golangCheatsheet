package main

import (
	"fmt"
)

type person struct {
	name     string
	last     string
	iceCream []string
}

func main() {
	p1 := person{
		name:     "Dennis",
		last:     "Ritchie",
		iceCream: []string{"Vanila", "Chocolate", "Coke"},
	}
	p2 := person{
		name:     "richard",
		last:     "Stallman",
		iceCream: []string{"Strawberry", "Neapolitan", "Buttered Pecan", "cappuccino"},
	}

	fmt.Println(p1.name)
	fmt.Println(p1.last)
	for i, v := range p1.iceCream {
		fmt.Println(i, v)
	}

	fmt.Println(p2.name)
	fmt.Println(p2.last)
	for i, v := range p2.iceCream {
		fmt.Println(i, v)
	}
}
