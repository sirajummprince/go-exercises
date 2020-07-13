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

func odd(f func(x ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(x...)
	fmt.Println(s)

	os := odd(sum, x...)
	fmt.Println(os)

}
