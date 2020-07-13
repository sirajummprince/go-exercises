package main

import (
	"fmt"
)

const (
	w = 2017
	a = iota
	x = w + a
	y = x + a
	z = y + a
)

func main() {
	fmt.Println(w, x, y, z)
}
