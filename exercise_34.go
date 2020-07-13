package main

import (
	"fmt"
)

func main() {
	v := struct {
		category  string
		color     map[string]string
		fuel_type []string
	}{
		category: `sedan`,
		color: map[string]string{
			`color1`: `white`,
			`color2`: `blue`,
			`color3`: `red`,
		},
		fuel_type: []string{
			`diesel`,
			`octane`,
			`petrol`,
		},
	}
	fmt.Println(v)
	for i, j := range v.color {
		fmt.Println(i, j)
	}
	for i, j := range v.fuel_type {
		fmt.Println(i, j)
	}
}
