package main

import (
	"encoding/json"
	"fmt"
)

type wizard struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {

	s := `[{"First":"Albus", "Last":"Dumbledore", "Age":116, "Sayings":["Don't pity the dead harry pity the living"]},{"First":"Minerva", "Last":"McGonagall", "Age":100, "Sayings":["I always wanted to use that spell"]},{"First":"Tom", "Last":"Riddle", "Age":72, "Sayings":["I am lord voldemort"]}]`
	fmt.Println(s)
	fmt.Println("------------------------")

	var wizards []wizard
	err := json.Unmarshal([]byte(s), &wizards)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wizards)

	for i, wizard := range wizards {
		fmt.Println("Wizard #", i)
		fmt.Println("\t", wizard.First, wizard.Last, wizard.Age)
		for _, saying := range wizard.Sayings {
			fmt.Println("\t", saying)
		}
	}

}
