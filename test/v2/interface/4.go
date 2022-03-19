package main

import "fmt"

func funcOne(i interface{}) {
	fmt.Printf("Type of i = %T and value = %v\n", i, i)
}

func main() {
	x := "abcdefghijklmnopq"
	funcOne(x)
	y := 100
	funcOne(y)
	s1 := struct {
		nameX string
	}{
		nameX: "XYZ",
	}
	funcOne(s1)

}
