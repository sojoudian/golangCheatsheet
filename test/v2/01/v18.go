package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4}
	fmt.Println(len(a))

	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("%dth  element is %d\n", i, a[i])
	// }

	for i, v := range a {
		fmt.Printf("%dth element is %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

}
