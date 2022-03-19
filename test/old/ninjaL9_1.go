package main

import "fmt"

func main() {
	go func() {
		fmt.Println("1")
	}()
	go func() {
		fmt.Println("2")
	}()

	fmt.Println("The program is about to exit")
}
