package main

import (
	"fmt"
)

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, y ...int) int {
	var y1 []int
	for _, v := range y {
		if v%2 == 0 {
			y1 = append(y1, v)
		}
	}
	return f(y1...)
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := sum(x...)
	fmt.Println(s)

	e := even(sum, x...)
	fmt.Println(e)
}
