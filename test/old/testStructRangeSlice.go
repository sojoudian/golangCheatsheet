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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(m)

	for j, k := range m {
		fmt.Println(j)
		fmt.Println(k.name)
		fmt.Println(k.last)
		for m, n := range k.iceCream {
			fmt.Println(m, n)
		}
		fmt.Println("--------------")

	}

}
