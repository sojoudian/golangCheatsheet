package main

import (
	"fmt"
	"time"
)

func hOne() {
	fmt.Println("hi")
}

func main() {
	go hOne()
	time.Sleep(2 * time.Second)
	fmt.Println("main function")
}
