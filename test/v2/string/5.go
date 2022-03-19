package main

import "fmt"

// func fn(s string) string {
// 	s[0] = '.'
// 	return s
// }

func fnCh(r []rune) string { // r is slice of rune
	r[0] = '.'
	result := string(r)
	return result
}

func main() {
	s := "hello"
	//fmt.Println(fn(s)) it will fail the program
	fmt.Println(fnCh([]rune(s)))
	//to check if the original array is effected or not
	fmt.Println(s)
}
