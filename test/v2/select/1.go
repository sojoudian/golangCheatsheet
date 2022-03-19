package main

import (
	"fmt"
	"time"
)

func xOne(channel chan string) {
	time.Sleep(6 * time.Second)
	channel <- "from xOne"
}

func xTwo(channel chan string) {
	time.Sleep(3 * time.Second)
	channel <- "from xTwo"
}

func main() {
	oOne := make(chan string)
	oTwo := make(chan string)

	go xOne(oOne)
	go xTwo(oTwo)
	select {
	case a1 := <-oOne:
		fmt.Println(a1)
	case a2 := <-oTwo:
		fmt.Println(a2)
	}
}
