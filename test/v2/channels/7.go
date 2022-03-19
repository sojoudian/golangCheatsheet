package main

import "fmt"

func pFunc(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func main() {
	channel := make(chan int)
	go pFunc(channel)
	for {
		x, y := <-channel
		if y == false {
			break
		}
		fmt.Println("Values and result", x, y)
	}
}
