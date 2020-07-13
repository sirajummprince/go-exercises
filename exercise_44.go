package main

import (
	"fmt"
)

func bar() func() int {
	return func() int {
		x := 119068
		return x
	}
}

func main() {
	y := bar()
	fmt.Println(y())
}
