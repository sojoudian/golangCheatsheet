package main

import "fmt"

func main() {
	channel := make(chan string, 3)
	channel <- "Albus"
	channel <- "Minerva"
	fmt.Println("channel capacity: ", cap(channel))
	fmt.Println("channel capacity: ", len(channel))
	// channel <- "Tom Riddle"

	fmt.Println("read value", <-channel)
	fmt.Println("channel length: ", len(channel))
}
