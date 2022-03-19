package main

import (
	"errors"
	"log"
)

//var errorNew := errors.New("This is the error message for  square root of negative numbers")

func main() {
	_, err := square(-10)
	if err != nil {
		log.Fatal(err)
	}
}

func square(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("This is the error message for  square root of negative numbers")
	}
	return 116, nil
}
