package main

import (
	"fmt"
)

type person struct {
	first_name          string
	last_name           string
	favourite_ice_cream []string
}

func main() {
	p1 := person{
		first_name:          `James`,
		last_name:           `Bond`,
		favourite_ice_cream: []string{`Vanilla`, `Mixed Fruit`},
	}
	p2 := person{
		first_name:          `Miss`,
		last_name:           `Moneypenny`,
		favourite_ice_cream: []string{`Chocolate`, `Four Seasons`},
	}
	fmt.Println(p1.first_name, p1.last_name, p1.favourite_ice_cream)
	fmt.Println(p2.first_name, p2.last_name, p2.favourite_ice_cream)

	fmt.Println(p1.first_name, p1.last_name)
	for i, v1 := range p1.favourite_ice_cream {
		fmt.Println(i, v1)
	}
	fmt.Println(p2.first_name, p2.last_name)
	for j, v2 := range p2.favourite_ice_cream {
		fmt.Println(j, v2)
	}
}
