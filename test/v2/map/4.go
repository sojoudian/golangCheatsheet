package main

import "fmt"

func main() {
	x := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for key, val := range x {
		fmt.Printf("x[%s]= %d\n", key, val)
	}
	fmt.Println("################################################################")
	delete(x, "a")
	for key, val := range x {
		fmt.Printf("x[%s]= %d\n", key, val)
	}

	x["b"] = 0
	fmt.Println("################################################################")
	for key, val := range x {
		fmt.Printf("x[%s]= %d\n", key, val)
	}
	fmt.Println("the lent of the x map is: ", len(x))
	fmt.Println("################################################################")
	y := x
	y["b"] = 0
	fmt.Println("value of y[b]", y["b"])
	fmt.Println(y)
}
