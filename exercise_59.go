package main

import (
	"fmt"
	"runtime"
	"sync"
)

const gs = 100

var counter int = 0
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Counter:\t", counter)
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
	fmt.Println("\n")

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//runtime.Gosched()
			v++
			counter = v
			fmt.Println("Counter:\t", counter)
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go routines:\t", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("\n")
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
}

//To check `race` condition run `$ go run -race exercise_59.go`
