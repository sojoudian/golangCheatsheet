package main

import "fmt"

func main() {
	defer first()
	second()
}

func first() {
	fmt.Println("Hello from First")
}

func second() {
	fmt.Println("Hello from Second")
}
