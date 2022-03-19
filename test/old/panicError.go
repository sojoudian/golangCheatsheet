package main

import (
	"os"
)

func main() {
	_, err := os.Open("notExisted.txt")
	if err != nil {
		panic(err)
	}
}
