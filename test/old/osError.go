package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	r := strings.NewReader("Hacker News")

	io.Copy(f, r)
}
