package main

import (
	"fmt"
	"sync"
	"time"
)

func x(i int, ab *sync.WaitGroup) {
	fmt.Println("go routine", i, "begins")
	time.Sleep(2 * time.Second)
	fmt.Printf("go routine %d ended\n", i)
	ab.Done()
}

func main() {
	v := 3
	var ab sync.WaitGroup
	for i := 0; i < v; i++ {
		ab.Add(1)
		go x(i, &ab)
	}
	ab.Wait()
	fmt.Println("all go routines have completed their execution")
}
