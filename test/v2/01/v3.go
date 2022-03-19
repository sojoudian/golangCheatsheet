package main

import "fmt"

var x int = 128

func main() {
	fmt.Println(x)
	//var y int = 256
	//y can not be used in function fnc because it is inside function main()
	fnc()

}

func fnc() {
	fmt.Println("from fnc", x)
	//fmt.Println("from fnc", y)
}
