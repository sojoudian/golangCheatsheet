package main

import "fmt"

func main() {
	var x, y, z int
	x, y = 75, 25
	z = x & y
	fmt.Println(z)

	var m, n int = 50, 60
	m += n
	fmt.Println(m)

	x -= n
	fmt.Println(x)
	x *= y
	fmt.Println(x)

	y = 10
	x = 12
	y %= x
	fmt.Println(y)

	y = 100
	x = 100
	y /= x
	fmt.Println(y)

}
