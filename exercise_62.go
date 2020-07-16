package main

import (
	"fmt"
)

func main() {
	//anonymous func
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	//buffered channel
	//no need of anonymous func
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}
