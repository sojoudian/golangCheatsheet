package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, value := range numbers {
		fmt.Println("Index: ", index, "Value: ", value)
	}
}
