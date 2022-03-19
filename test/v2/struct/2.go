package main

import "fmt"

func main() {
	a := struct {
		fruits, vegetables string
	}{
		fruits:     "Apple",
		vegetables: "Cauliflower",
	}
	fmt.Println("a", a)
}
