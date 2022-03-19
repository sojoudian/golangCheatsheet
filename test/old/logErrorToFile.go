package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.log")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	ff, err := os.Open("notExisteFile.txt")
	if err != nil {
		log.Println("Error occurred", err)
	}
	defer ff.Close()
	fmt.Println("Check the log.log for more information")
}
