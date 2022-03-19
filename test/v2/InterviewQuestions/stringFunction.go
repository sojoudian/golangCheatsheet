package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("maziar", "mazi"))
	fmt.Println(strings.Contains("maziar", "Riddle"))
	fmt.Println(strings.HasPrefix("maziar", "m"))
	fmt.Println(strings.HasPrefix("maziar", "Riddle"))
	fmt.Println(strings.HasSuffix("maziar", "r"))
	fmt.Println(strings.Index("maziar", "i"))
	fmt.Println(strings.Index("maziar", "m"))
	fmt.Println(strings.Index("maziar", "s"))
	fmt.Println(strings.Index("maziar", "M"))
	fmt.Println(strings.Join([]string{"Tom", "Marvolo", "Riddle"}, " "))
	fmt.Println(strings.Join([]string{"Tom", "Marvolo", "Riddle"}, "-"))
	fmt.Println(strings.Repeat("Marvolo ", 3))
	fmt.Println(strings.Replace("Marvolo", "o", "x", 2))
	fmt.Println(strings.Replace("Marvolo", "o", "x", 1))
	fmt.Println(strings.Replace("MarvoloMarvolo", "Marvolo", "Harry", 1))
	fmt.Println(strings.Replace("MarvoloMarvolo", "Marvolo", "Harry", 2))
	fmt.Println(strings.Split("Marvolo-Marvolo-Marvolo", "-"))
	fmt.Println(strings.Split("Tom-Marvolo-Riddle", "-"))
	fmt.Println(strings.ToLower("LORD VOLDEMORT"))
	fmt.Println(strings.ToUpper("lord voldemort"))

}
