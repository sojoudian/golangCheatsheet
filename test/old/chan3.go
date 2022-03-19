package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 116
	//c <- 11

	fmt.Println(<-c)
}
