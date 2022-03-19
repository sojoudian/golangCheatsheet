package main

import "fmt"

func main() {
	x := map[string]int{
		"a": 1,
		"b": 2,
	}

	key := "a"
	fmt.Println(x)
	fmt.Println("valu of x[key] is", x[key])
	fmt.Println("lets try a not existed key, valu of x[test]: ", x["test"])

	// z := map[string]int{
	// 	"a":    1,
	// 	"b":    2,
	// 	"test": 3,
	// }

	val, ok := x[key]
	if ok == true {
		fmt.Println("The key", key, " is existed and the value is", val)
	} else {
		fmt.Println("The key is not present", val)
	}

}
