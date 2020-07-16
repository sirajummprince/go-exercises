package main

import (
	"fmt"
)

func main() {
	//first code
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)
	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	//second code
	cr := make(chan int)

	go func() {
		cr <- 43
	}()
	fmt.Println(<-cr)
	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cr)
}
