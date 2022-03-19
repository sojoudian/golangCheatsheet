package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old")
}
func main() {
	w1 := person{
		first: "Albus",
		last:  "Dumbledore",
		age:   116,
	}
	w1.speak()
}
