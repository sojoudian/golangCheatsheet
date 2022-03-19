package main

import "fmt"

func main() {
	ve(1, "a")
}

func ve(x int, s ...string) {
	fmt.Println(x, s)
}
