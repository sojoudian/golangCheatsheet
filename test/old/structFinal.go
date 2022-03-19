package main

import (
	"fmt"
)

func main() {
	s := struct {
		name    string
		last    string
		contact map[string]int
		drink   []string
	}{
		name: "Richard",
		last: "Stallman",
		contact: map[string]int{
			"Ken":   111,
			"Denis": 222,
		},
		drink: []string{
			"Water",
			"Sosa",
		},
	}
	fmt.Println(s)

	for i, j := range s.contact {
		fmt.Println(i, j)
	}
	for i, j := range s.drink {
		fmt.Println(i, j)
	}

}
