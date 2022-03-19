package main

import "fmt"

func fn(num [4]int) {
	num[0] = 12
	fmt.Println("from fn", num)
	fmt.Printf("arr num %d", num[0])
}
func main() {
	var a [4]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	//a[4] = 5
	fmt.Println("arr a", a)

	b := [4]int{1, 2, 3, 4}
	fmt.Println("arr b", b)

	c := [4]int{1}
	fmt.Println("arr c", c)
	d := [...]int{1, 2}
	fmt.Println("arr d", d)

	e := [3]int{1, 2, 3}
	fmt.Println("arr e", e)

	// var f [5]int
	// f = e
	// fmt.Println("arr f", f)

	g := [...]string{"a", "b", "c"}
	h := g
	h[0] = "x"
	fmt.Println("arr h", h)
	fmt.Println("arr g", g)

	num := [...]int{1, 2, 3, 4}
	fmt.Println("num in main()", num)
	fn(num)
	fmt.Println("back in function main()")
	fmt.Println("arr from fn num", num)
}
