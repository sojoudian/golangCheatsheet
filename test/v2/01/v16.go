package main

import (
	"fmt"
	"time"
)

func rn() {
	if r := recover(); r != nil {
		fmt.Println("recovered", r)
	}
}

func x() {
	defer rn()
	fmt.Println("in x")
	//this program will not recover anything in func rn because
	//go y() will create another goroutine and in this will not working
	//in the other go routin
	/////////////////////////////////////////////////////////////////////////////go y()
	//but if we change
	//go y()
	//to
	//y()
	//it will work prefectly
	y()
	time.Sleep(1 * time.Second)
}

func y() {
	fmt.Println("in y")
	panic("panicking in y")
}
func main() {
	x()
	fmt.Println("returning back to main()")

}
