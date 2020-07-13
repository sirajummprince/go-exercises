package main

import (
	"fmt"
)

func main() {
	x := 119068
	if x == 119067 {
		fmt.Printf("x is %v", x)
	} else if x == 119066 {
		fmt.Printf("x is %v ", x)
	} else {
		fmt.Println("if, else if, else conditionals work")
	}
}
