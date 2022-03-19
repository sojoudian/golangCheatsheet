package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("Channel is nil")
		fmt.Printf("%T\n", a)
		a = make(chan int)
		fmt.Printf("%T\n", a)
	}
}
