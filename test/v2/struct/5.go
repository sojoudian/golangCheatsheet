package main

import "fmt"

type Diner struct {
	Food     string
	Calories int
}

type Menu struct {
	Fruits string
	items  Diner
}

func main() {
	var b Menu
	b.Fruits = "Apple"
	b.items.Food = "Steak"
	b.items.Calories = 10
	fmt.Println(b)

}
