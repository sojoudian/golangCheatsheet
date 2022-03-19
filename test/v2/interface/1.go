package main

import "fmt"

type X interface {
	myfunction() []rune
}

type Varstr string //variable

func (str Varstr) myfunction() []rune {
	var xy []rune
	for _, r := range str {
		if r == 'a' || r == 'i' || r == 'e' || r == 'o' || r == 'u' {
			xy = append(xy, r)
		}
	}
	return xy
}

func main() {
	n := Varstr("abcdefghijklmnopq")
	var m X
	m = n
	fmt.Println("Display:", m.myfunction())
}
