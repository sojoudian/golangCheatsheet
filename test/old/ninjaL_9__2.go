package main

import "fmt"

type person struct {
	First string
}

func (p *person) speak() {
	fmt.Println("Hello")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	w1 := person{
		First: "Albous",
	}
	// THIS WILL NOT WORK
	//saySomething(w1)
	saySomething(&w1)

	w1.speak()

}
