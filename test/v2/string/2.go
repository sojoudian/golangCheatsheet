package main

import "fmt"

func main() {
	s := []byte{0x43, 0x61} //Hexadecimal
	d := []byte{67, 97}     //ASCII
	dStr := string(d)

	str := string(s)
	fmt.Println(str)
	fmt.Println(dStr)
}
