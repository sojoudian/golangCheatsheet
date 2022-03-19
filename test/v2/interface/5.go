package main

import "fmt"

func x(i interface{}) {
	y := i.(int)
	fmt.Println(y)
}

func funcX(i interface{}) {
	b, y := i.(int)
	fmt.Println(b, y)
}

func main() {
	var g interface{} = 64
	x(g)
	// var h interface{} = "A"
	// x(h)
	//it is not working
	var h interface{} = "A"
	funcX(h)
	var i interface{} = 64
	funcX(i)

}
