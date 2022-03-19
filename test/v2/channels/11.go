package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Albus"
	channel <- "Minerva"
	channel <- "Tom Riddle"
	//wont work
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
