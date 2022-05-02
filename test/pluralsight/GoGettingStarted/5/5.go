package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	slice := arr[:]

	fmt.Println(arr, slice)

	arr[1] = 4
	slice[2] = 9
	fmt.Println(arr, slice)

}
