package main

import "fmt"

func main() {
	var firstName *string = new(string)
	*firstName = "Rob"
	fmt.Println(firstName)
	fmt.Println(*firstName)
	fmt.Println(&firstName)

	lastName := "Pike"
	fmt.Println(lastName)
	pointer := (&lastName)
	fmt.Println(pointer, *pointer)

	lastName = "Thompson"
	fmt.Println(pointer, *pointer)
}
