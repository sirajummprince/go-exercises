package main

import (
	"fmt"
	"sync"
)

var gs int = 10
var wg sync.WaitGroup

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i <= gs; i++ {
			wg.Add(1)
			go func() {
				for j := 0; j <= 10; j++ {
					c <- j
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

}
