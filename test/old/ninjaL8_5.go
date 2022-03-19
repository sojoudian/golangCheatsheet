package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type wizard struct {
	Name    string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []wizard

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []wizard

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	w1 := wizard{
		Name:    "Albus",
		Last:    "Dumbledore",
		Age:     116,
		Sayings: []string{"Don't pity the dead harry pity the living"},
	}
	w2 := wizard{
		Name:    "Minerva",
		Last:    "McGonagall",
		Age:     100,
		Sayings: []string{"I always wanted to use that spell"},
	}
	w3 := wizard{
		Name:    "Tom",
		Last:    "Riddle",
		Age:     72,
		Sayings: []string{"I am lord voldemort"},
	}
	wizards := []wizard{w1, w2, w3}
	fmt.Println(wizards)

	err := json.NewEncoder(os.Stdout).Encode(wizards)
	if err != nil {
		fmt.Println("An error has occurred")
	}

	sort.Sort(ByAge(wizards))
	for _, w := range wizards {
		fmt.Println(w.Name, w.Last, w.Age)
		for _, v := range w.Sayings {
			fmt.Println("\t", v)
		}
	}
	fmt.Println("--------------")
	sort.Sort(ByLast(wizards))
	for _, w := range wizards {
		fmt.Println(w.Name, w.Last, w.Age)
		for _, v := range w.Sayings {
			fmt.Println("\t", v)
		}
	}
}
