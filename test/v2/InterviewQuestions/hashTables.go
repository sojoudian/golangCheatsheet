package main

import "fmt"

//How map can be used as hash table
//masps are unordered collections and there is no way
//to predict the order in which the key value pairs will be returned
// it means everytime you run the program, every iteration
// over the map could return different order

func main() {
	var cx map[int]string
	cx = make(map[int]string)
	cx[1] = "a"
	cx[2] = "b"
	cx[3] = "c"
	for i, j := range cx {
		fmt.Println("key: ", i, "value: ", j)
	}
}
