package main

import "fmt"

func main() {
	ru := []rune{0x0053, 0x0065, 0x00f1}
	str := string(ru)
	fmt.Println(str)
}
