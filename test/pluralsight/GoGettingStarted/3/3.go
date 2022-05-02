package main

import "fmt"

const (
	pi     = 3.1415
	first  = 1
	second = "second"
	third  = iota
	fourth = iota
	fifth  = iota + 6
	sixth  = 2 << iota
	seventh
)

const (
	one = iota
	two
)

func main() {
	fmt.Println(pi, first, second, third, fourth, fifth, sixth, seventh, one, two)
}
