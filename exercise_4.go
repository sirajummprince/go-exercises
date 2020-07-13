package main

import (
	"fmt"
)

type prince int

var x prince

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
