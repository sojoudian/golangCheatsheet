package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("notExisted.txt")
	if err != nil {
		log.Println("error occurred opening the file notExisted.txt", err)
	}
}
