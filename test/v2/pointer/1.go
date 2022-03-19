package main

import "fmt"

func fn(p *int) {
	*p = 77
}

func fp(p *int) {
	*p = 88
}

func main() {
	b := 100
	var a *int = &b //integer pointer to address of b
	fmt.Println("a is: ", a)
	fmt.Println("b is: ", &b)
	fmt.Printf("%T\n ", a)
	fmt.Printf("%d\n", *a)
	//fmt.Println("value of b through a ", *a)

	fmt.Println("before: ", b)
	*a++
	fmt.Print("after: ", b)
	fn(a)
	fmt.Println("a after fn function a is: ", a)
	fn(a)
	fmt.Println("b after fn function a is: ", b)

	fn(&b)
	fmt.Println("b after fn function fn is: ", b)

	fp(&b)
	fmt.Println("b after fn function fp is: ", b)

	var c *int
	if c == nil {
		fmt.Println("\n c is", c)
	}
}
