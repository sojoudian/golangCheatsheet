package main

import "fmt"

func main() {
	cs := make(chan int)

	go func() {
		cs <- 116
	}()
	fmt.Println(<-cs)
	fmt.Printf("---------\n")
	fmt.Printf("cs\t%T\n", cs)
}
