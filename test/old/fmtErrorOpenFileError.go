package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("notExisted.txt")
	if err != nil {
		fmt.Println("error occurred opening the file notExisted.txt", err)
	}
}
