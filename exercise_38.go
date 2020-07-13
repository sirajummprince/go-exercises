package main

import (
	"fmt"
)

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := foo(x...)
	z := bar(x)
	fmt.Println("With unfurling", y)
	fmt.Println("With slice", z)
}
