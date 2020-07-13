package main

import (
	"fmt"
)

type persion struct {
	first_name          string
	last_name           string
	favourite_ice_cream []string
}

func main() {
	p1 := persion{
		first_name:          `James`,
		last_name:           `Bond`,
		favourite_ice_cream: []string{`Vanilla`, `Mixed Fruit`},
	}
	p2 := persion{
		first_name:          `Miss`,
		last_name:           `Moneypenny`,
		favourite_ice_cream: []string{`Chocolate`, `Passion Fruit`},
	}
	m := map[string]persion{
		p1.last_name: p1,
		p2.last_name: p2,
	}
	for i, v := range m {
		fmt.Println(i)
		for j, k := range v.favourite_ice_cream {
			fmt.Println(j, k)
		}
	}
}
