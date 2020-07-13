package main

import (
	"fmt"
)

const (
	x int = 2
	y     = x << 1
)

func main() {
	fmt.Println("\tdecimal\tbinary\thex")
	fmt.Printf("x\t%d\t%b\t%#x\n", x, x, x)

	fmt.Println("\tdecimal\tbinary\thex")
	fmt.Printf("y\t%d\t%b\t%#x", y, y, y)
}
