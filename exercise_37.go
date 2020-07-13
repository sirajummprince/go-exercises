package main

import (
	"fmt"
)

func foo(i int) int {
	return i
}
func bar(i int, j string) (int, string) {
	return i, j
}

func main() {
	x := foo(119068)
	z, y := bar(104446, `Prince`)
	fmt.Println(x, y, z)
}
