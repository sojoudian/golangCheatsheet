package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type wizard struct {
	Name    string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	w1 := wizard{
		Name:    "Albus",
		Last:    "Dumbledore",
		Age:     116,
		Sayings: []string{"Don't pity the dead harry pity the living"},
	}
	w2 := wizard{
		Name:    "Minerva",
		Last:    "McGonagall",
		Age:     100,
		Sayings: []string{"I always wanted to use that spell"},
	}
	w3 := wizard{
		Name:    "Tom",
		Last:    "Riddle",
		Age:     72,
		Sayings: []string{"I am lord voldemort"},
	}
	wizards := []wizard{w1, w2, w3}
	fmt.Println(wizards)

	err := json.NewEncoder(os.Stdout).Encode(wizards)
	if err != nil {
		fmt.Println("An error has occurred")
	}
}
