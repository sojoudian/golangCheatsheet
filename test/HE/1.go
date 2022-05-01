package main

import (
	"fmt"
	"log"
)

const (
	maxLength  = 20
	YES_ANSWER = "Yes"
	NO_ANSWER  = "No"
)

func main() {
	var input string
	fmt.Print("Enter the string here: ")
	fmt.Scanf("%s", &input)

	if len(input) > maxLength {
		log.Panic("The maximum length of Input must be 20 characters")

	} else {
		fmt.Println("Hello", input)
		var z int
		var o int
		for _, c := range input {
			if string(c) == "z" {
				z++
			} else {
				o++
			}
		}
		if o/z != 2 {
			fmt.Println(NO_ANSWER)
		} else {
			fmt.Println(YES_ANSWER)
		}

	}
}
