package main

import "fmt"

func main() {
	user := map[string]string{
		"name":  "Albus",
		"email": "albus@hg.wiz",
		"role":  "Head master",
	}
	user["age"] = "30"

	for key, val := range user {
		fmt.Println("Key: ", key, "Value: ", val)
	}
}
