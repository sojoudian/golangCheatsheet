package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4}
	fmt.Println(slc)
	lastElement := slc[len(slc)-1]
	fmt.Println("The last element is: ", lastElement)
	firstElement := slc[0]
	fmt.Println("The First element is: ", firstElement)
	removedElement := slc[:len(slc)-1]
	fmt.Println("Slice, after removed element: ", removedElement)
	afterFirstElement := slc[1:]
	fmt.Println("new Slice: ", afterFirstElement)
}
