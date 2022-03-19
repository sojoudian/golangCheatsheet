package main

import "fmt"

func main() {
	ve()
	ve("abc", "efg", "ghi")
}

func ve(s ...string) {
	// fmt.Println(s[0])
	// fmt.Println(s[1])
	// fmt.Println(s[2])
	fmt.Println(s)
}
