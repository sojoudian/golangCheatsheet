package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begins CPU", runtime.NumCPU())
	fmt.Println("begins goRoutines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("1")
		wg.Done()
	}()
	go func() {
		fmt.Println("2")
		wg.Done()
	}()
	fmt.Println("mid CPU", runtime.NumCPU())
	fmt.Println("mid goRoutines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("The program is about to exit")
	fmt.Println("end CPU", runtime.NumCPU())
	fmt.Println("end goRoutines", runtime.NumGoroutine())
}
