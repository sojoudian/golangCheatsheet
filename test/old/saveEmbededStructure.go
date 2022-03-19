package main

import (
	"fmt"
)

type vehicle struct {
	door  int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			door:  2,
			color: "red",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			door:  4,
			color: "Black",
		},
		luxury: true,
	}
	fmt.Println(t, s)
}
