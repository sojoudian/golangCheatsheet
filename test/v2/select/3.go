package main

import (
	"fmt"
	"time"
)

func pOne(channel chan string) {
	time.Sleep(10500 * time.Millisecond)
	channel <- "pOne data has been written"
}

func main() {
	channel := make(chan string)
	go pOne(channel)
	for {
		time.Sleep(1 * time.Second)
		select {
		case varx := <-channel:
			fmt.Println("received value", varx)
			return
		default:
			fmt.Println("value is not received")

		}
	}
}
