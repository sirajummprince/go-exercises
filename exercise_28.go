package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:     []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`monepenny_miss`: []string{`Jame Bond`, `Literature`, `Computer Science`},
		`no_dr`:          []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for i, v := range m {
		fmt.Printf("Name: %v\n", i)
		fmt.Println("Favourites:\n")
		for j, k := range v {
			fmt.Println(j, k)
		}
		fmt.Println("\n")
	}
}
