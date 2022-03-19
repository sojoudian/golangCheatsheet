package main

import (
	"fmt"
)

func main() {
	bd := 1989
	for {
		if bd > 2021 {
			break
		}
		fmt.Println(bd)
		bd++
	}

}
