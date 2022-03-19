package main

import (
	"encoding/json"
	"fmt"
)

type wizard struct {
	First string
	Age   int
}

func main() {
	w1 := wizard{
		First: "Albus",
		Age:   116,
	}
	w2 := wizard{
		First: "Minerva",
		Age:   100,
	}
	w3 := wizard{
		First: "Tom",
		Age:   72,
	}

	wizards := []wizard{w1, w2, w3}

	fmt.Println(wizards)
	bs, err := json.Marshal(wizards)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
