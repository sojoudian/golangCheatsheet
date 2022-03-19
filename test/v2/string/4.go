package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := []byte{0x43, 0x61} //Hexadecimal
	d := []byte{67, 97}     //ASCII
	dStr := string(d)

	str := string(s)
	fmt.Println(str)
	fmt.Println(dStr)
	ru := []rune{0x0053, 0x0065, 0x00f1}
	ruStr := string(ru)
	fmt.Println(ruStr)
	fmt.Println(ruStr, utf8.RuneCountInString((ruStr)))
	fmt.Println(str, utf8.RuneCountInString((str)))
}
