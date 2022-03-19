// Maps can be constructed using the usual composite literal syntax with colon-separated key-value pairs, so it's easy to build them during initialization.

// var timeZone = map[string]int{
//     "UTC":  0*60*60,
//     "EST": -5*60*60,
//     "CST": -6*60*60,
//     "MST": -7*60*60,
//     "PST": -8*60*60,
// }
// like ./sliceCompositeLiteral.go:8:

package main

import (
	"fmt"
)

func main() {
	//u1 := []string{"James", "Bond", "Chocolate", "martini"}
	m := map[string]int{
		"James": 32,
		"penny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["sheldon"])

	//Golang comma ok idiom
	v, ok := m["sheldon"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["sheldon"]; ok {
		fmt.Println("if print", v)
	} else {
		fmt.Println("the key doesn't exist")
	}

	m["Amy"] = 25
	fmt.Println(m)
	m["Alex"] = 24
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

}
