package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	ip := net.ParseIP(os.Args[1])

	if ip != nil {
		fmt.Printf("%v OK\n", ip)
	} else {
		fmt.Println("Bad IP address: ", ip)
	}
}
