package main

import "fmt"

func main() {
	f := foo()
	i, s := bar()
	fmt.Println(f, i, s)
}

func foo() int {
	return 110
}

func bar() (int, string) {
	return 1984, "The Big Brother is Watching You"
}
