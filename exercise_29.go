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
	m[`q_m`] = []string{`007`, `James Bond`, `Moneypenny`}
	for i, v := range m {
		fmt.Println("Index:", i)
		for j, k := range v {
			fmt.Println(j, k)
		}
		fmt.Println("\n")
	}

}
