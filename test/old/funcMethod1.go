package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type wizard struct {
	person
	itb bool
}

//func (r receiver) identifire(parameters) (return(s))

func (m wizard) spell() {
	fmt.Println("I am, ", m.first, m.last)
}

func main() {
	w1 := wizard{
		person: person{
			first: "Albus",
			last:  "Dumbledore",
		},
		itb: true,
	}

	w2 := wizard{
		person: person{
			first: "Minerva",
			last:  "McGonagall",
		},
		itb: true,
	}

	fmt.Println(w1, w2)
	w1.spell()
	w2.spell()
}
