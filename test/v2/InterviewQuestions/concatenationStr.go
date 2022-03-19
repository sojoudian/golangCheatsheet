package main

import "fmt"

func main() {
	fmt.Println("1st string")
	var f string
	fmt.Scanln(&f)
	fmt.Println("2nd string")
	var s string
	fmt.Scanln(&s)
	fmt.Println(f + s)
}
