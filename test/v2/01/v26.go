package main

import "fmt"

func main() {
	a := [][]string{
		{"a", "b"},
		{"c", "d"},
	}

	for _, v := range a {
		for _, j := range v {
			fmt.Printf("%s", j)
		}
		fmt.Println("\n")
	}
}
