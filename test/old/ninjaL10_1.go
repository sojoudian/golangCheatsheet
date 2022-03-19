package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 116
	}()

	fmt.Println(<-c)
}
