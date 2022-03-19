package main

import "fmt"

func main() {
	fmt.Println("Please enter 3 numbers: ")
	var f int
	var s int
	var t int
	fmt.Scanln(&f)
	fmt.Scanln(&s)
	fmt.Scanln(&t)
	if f >= s && f >= t {
		fmt.Println("The largest number is: ", f)
	} else if s >= f && s >= t {
		fmt.Println("The largest number is: ", s)
	} else if t >= f && t >= s {
		fmt.Println("The largest number is: ", t)
	}
}
