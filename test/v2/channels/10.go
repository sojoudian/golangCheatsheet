package main

import (
	"fmt"
	"time"
)

func cCHan(channel chan int) {
	for y := 0; y < 5; y++ {
		channel <- y
		fmt.Println(y, "write to the channel")
	}
	close(channel)
}

func main() {
	channel := make(chan int, 2)
	go cCHan(channel)
	time.Sleep(2 * time.Second)
	for val := range channel {
		fmt.Println(val, "write to the channel cCHan")
		time.Sleep(2 * time.Second)
	}
}
