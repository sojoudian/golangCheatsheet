package main

func main() {
	a := make(chan int)
	a <- 100
	// thi program will not work
}
