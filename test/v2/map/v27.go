package main

import "fmt"

func main() {
	//x := make(map[string]int)
	var x map[string]int
	if x == nil {
		fmt.Println("the map named x is nill")
		x = make(map[string]int)
	}
}
