package main

import (
	"fmt"
	"math"
)

type circlr struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circlr) area() float64 {
	return math.Pi * c.radius * c.radius
}

func infos(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circlr{5}
	//infos(c)
	infos(&c)
	//fmt.Println(c.area())
}
