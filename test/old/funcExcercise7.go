package main

import "fmt"

var x int

var y int = 256

var g func() = func() {
	fmt.Println("g from outside main")
}

func main() {
	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// g from out side :
	g()
	//g from inside
	g = f
	g()
	fmt.Printf("This is G %T\n", g)
	fmt.Println("Done")

	fmt.Println(y)

}
