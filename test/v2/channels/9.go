package main

import "fmt"

//send and receiveto an unbuffered channelis locked
//buffered channel
//x := make(chan type, capacity)

func main() {
	channel := make(chan string, 2)
	channel <- "abc"
	channel <- "def"
	// channel <- "ghi"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
