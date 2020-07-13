package main

import (
	"fmt"
)

var f func()

func main() {
	x := func() {
		fmt.Println("foo is assigned to x and printed")
	}
	f = x
	f()
}
