package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	s2 := f()
	fmt.Println(s2)

	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)

	fmt.Println("run the function directly from print", x())
}

func foo() string {
	s := "Hello world"
	return s
}

func f() string {
	return "this is another method for returning a string"
}
func bar() func() int {
	return func() int {
		return 110
	}
}
