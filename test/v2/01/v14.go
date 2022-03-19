package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter 1 for junior, 2 for senior")
	fmt.Scanln(&a)
	switch a {
	case 1:
		fmt.Println("Junior Developer")
	case 2:
		fmt.Println("Senior Developer")
	default:
		//fmt.Printf("before %d", a, "\n")
		panic(fmt.Sprintf("Invalid %d", a))
		fmt.Printf("after %d", a)
	}
}
