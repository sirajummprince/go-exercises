package main

import (
	"fmt"
)

func main() {
	x1 := []string{"Go", "Python", "Javascript", "Julia"}
	x2 := []string{"C", "C++", "Java", "C#"}
	x := [][]string{x1, x2}
	for i, val1 := range x {
		fmt.Println(i, "\t")
		for j, val2 := range val1 {
			fmt.Println("\t", j, val2)
		}
	}
}
