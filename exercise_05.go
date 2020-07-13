package main

import "fmt"

type prince int

var x prince

var y int

func main() {
	fmt.Println(x)

	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
