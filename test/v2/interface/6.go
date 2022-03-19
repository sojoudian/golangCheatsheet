package main

import "fmt"

func funcOne(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("value of string is %s\n", i.(string))
	case int:
		fmt.Printf("value of int is %d\n", i.(int))
	default:
		fmt.Printf("does not belong to any of defined type\n")
	}
}

func main() {
	funcOne(123)
	funcOne("name")
	funcOne(3.14)
}
