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
	var a Menu
	a.Fruits = "Pomegranate"
	a.Food = "Steak" //b.items.Food = "Steak"
	a.Calories = "8" //b.items.Calories = 10
	fmt.Println(a)
}
