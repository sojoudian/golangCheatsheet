package main

import "fmt"

type Basket struct {
	fruits, vegetables string
	price              int
}

func main() {
	a := &Basket{"Apple", "Lettuce", 6}
	fmt.Println("a: ", a)
	fmt.Println("Fruites:", a.fruits)
	fmt.Println("Vegetables:", a.vegetables)
	(*a).fruits = "Banana"
	fmt.Println("Fruites:", (*a).fruits)
}
