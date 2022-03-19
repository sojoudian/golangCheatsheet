package main

import "fmt"

func main() {
	var m, n int = 20, 30
	var s int
	var a, b string = "abc", "def"
	var d, p, q, r int
	var str string
	s = m + n
	fmt.Println(s)

	str = a + b
	fmt.Println(str)

	d = m - n
	p = m * n
	q = m / n
	r = m % n
	fmt.Println(d, p, q, r)
}
