package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementer int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)

			fmt.Println("inside func", atomic.LoadInt64(&incrementer))

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Incrementer final value: ", incrementer)
}
