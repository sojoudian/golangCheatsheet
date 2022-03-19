package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"sheldon": "cooper",
		"penny":   "couco",
		"Emily":   "sweeney",
	}
	//fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	// fmt.Println("\n")

	y := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	y[`Missy`] = []string{`cooper`, `not genius`, `tall`}

	for i, v := range y {
		fmt.Println("this is  index", i)
		for j, k := range v {
			fmt.Println("\t", j, k)
		}
	}

}
