package main

import (
	"errors"
	"fmt"
	"log"
)

var errorNew = errors.New("This is the error message for  square root of negative numbers")

func main() {
	fmt.Printf("%T\n", errorNew)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errorNew
	}
	return 116, nil
}
