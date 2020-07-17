package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c

	fmt.Println("Before closing\t", v, ok)

	close(c)

	v, ok = <-c

	fmt.Println("After closing\t", v, ok)
}
