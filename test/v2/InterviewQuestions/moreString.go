package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Albus", "l"))
	fmt.Println(strings.ContainsAny("Albus", "a"))
	fmt.Println(strings.ContainsAny("Albus", "Aa"))
	fmt.Println(strings.ContainsAny("Albus", "tom"))
	fmt.Println(strings.Contains("Albus", "Al"))
	fmt.Println(strings.Contains("Albus", "Ab"))
	fmt.Println(strings.Count("AAAlbus", "A"))
	fmt.Println(strings.Count("AlbusAlbus", "Albus"))
	fmt.Println(strings.EqualFold("Albus", "albus"))
	fmt.Println(strings.EqualFold("Albus", "albusX"))
}
