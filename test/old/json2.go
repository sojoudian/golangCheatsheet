package main

import (
	"encoding/json"
	"fmt"
)

type wizard struct {
	Name string `json:"Name"`
	Last string `json:"Last"`
	Age  int    `json:"Age"`
}

func main() {
	s := `[{"Name":"Albus","Last":"Dumbledore","Age":116},{"Name":"Minerva","Last":"McGonagall","Age":100}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//phoenixOrder := []wizard{}
	var phoenixOrder []wizard

	err := json.Unmarshal(bs, &phoenixOrder)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", phoenixOrder)

	for i, v := range phoenixOrder {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.Name, v.Last, v.Age)
	}

}
