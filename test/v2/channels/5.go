package main

func x(xa chan<- int) {
	xa <- 10
}

func main() {
	xa := make(chan<- int)
	go x(xa)
	// fmt.Println(<-xa) // wont work and we need conversional channel

}
