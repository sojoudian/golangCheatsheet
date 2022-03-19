package main

import (
	"encoding/json"
	"fmt"
)

type wizard struct {
	Name string
	Last string
	Age  int
}

func main() {
	w1 := wizard{
		Name: "Albus",
		Last: "Dumbledore",
		Age:  116,
	}
	w2 := wizard{
		Name: "Minerva",
		Last: "McGonagall",
		Age:  100,
	}
	phoenixOrder := []wizard{
		w1,
		w2,
	}

	fmt.Println(phoenixOrder)
	bs, err := json.Marshal(phoenixOrder)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
