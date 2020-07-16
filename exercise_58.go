package main

import (
	"fmt"
	"runtime"
	"sync"
)

const gs = 100

var counter int = 0
var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Counter:\t", counter)
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
	fmt.Println("\n")

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("Counter:\t", counter)
			wg.Done()
		}()
		fmt.Println("Go routines:\t", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("\n")
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
}

//To check `race` condition run `$ go run -race exercise_58.go`
