package main

import "fmt"

func main() {
	cr := make(chan int)

	go func() {
		cr <- 116
	}()

	fmt.Println(<-cr)
	fmt.Printf("---------\n")
	fmt.Printf("cr\t%T\n", cr)

}
