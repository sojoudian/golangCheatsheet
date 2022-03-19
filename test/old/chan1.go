package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 116

	fmt.Println(<-c)
}

// this program is not working because channles requires the send and receive at same time
