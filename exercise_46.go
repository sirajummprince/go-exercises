package main

import (
	"fmt"
)

var x int = 20

func foo() {
	x := 2
	fmt.Println(x, "is enclosed in foo")
}

func main() {
	fmt.Println(x, "is enclosed in package")
	foo()
	x := 200
	fmt.Println(x, "is enclosed in main")
}
