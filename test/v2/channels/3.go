package main

import (
	"fmt"
	"time"
)

func fOne(d chan bool) {
	fmt.Println("funcOne")
	time.Sleep(5 * time.Second)
	fmt.Println("waking from the sleep in the fOne go routine")
	d <- true
}

func main() {
	d := make(chan bool)
	fmt.Println("Calling fOne >>")
	go fOne(d)
	<-d
	fmt.Println("main()")
}
