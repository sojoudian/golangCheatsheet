package main

import "fmt"

//deadlock

func main() {
	channel := make(chan string)
	select {
	case <-channel:
	default:
		fmt.Println("to avoid deadlock")
	}
}
