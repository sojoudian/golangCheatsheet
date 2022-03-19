package main

import "fmt"

func fn(p *[3]int) { //it accepts a pointer to 3 integers
	(*p)[0] = 0 // dereffrencing p
}

func f(p *[3]int) { //it accepts a pointer to 3 integers
	p[0] = 89 // dereffrencing p
}

func main() {
	a := [3]int{89, 90, 91}
	fn(&a)
	fmt.Println(a)

	f(&a)
	fmt.Println(a)

}
