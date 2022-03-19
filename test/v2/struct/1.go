package main

import "fmt"

type Basket struct {
	fruits, vegetables string
	count              int
}

func main() {
	a := Basket{
		fruits:     "Apple",
		vegetables: "Carrot",
		//count:      10,
	}

	b := Basket{
		"Pomegranate", "Lettuce", 8,
	}
	fmt.Println("a", a)
	fmt.Println("b", b)

	var c Basket
	fmt.Println("c:", c)

	//accessing the individual fields of the struct:
	fmt.Println("accessing individual fields of struct a", a.fruits, a.vegetables)

	var d Basket
	fmt.Println("d: ", d)
	d.fruits = "Grapes"
	d.vegetables = "Lettuce"
	d.count = 2
	fmt.Println("d: ", d)
}
