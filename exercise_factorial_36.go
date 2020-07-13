package main

import (
	"fmt"
)

func factorial(n int) int {
	f := 1
	for ; n > 0; n-- {
		f *= n
	}
	return f
}

func main() {
	x := factorial(4)
	fmt.Println(x)
}
