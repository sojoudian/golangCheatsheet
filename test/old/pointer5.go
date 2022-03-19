package main

import "fmt"

type person struct {
	name string
	// last  string
	// age int
}

func main() {
	w1 := person{
		name: "Albus Dumbledore",
	}

	fmt.Println(w1)
	changeMe(&w1)
	fmt.Println(w1)

}

func changeMe(p *person) {
	p.name = "Tom Riddle"
	//both are the same
	(*p).first = "Tom Riddle"

}
