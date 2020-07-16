package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

const gs = 100

var counter int64
var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Counter:\t", counter)
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
	fmt.Println("\n")

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			v := atomic.LoadInt64(&counter)
			fmt.Println("Counter:\t", v)
			wg.Done()
		}()
		fmt.Println("Go routines:\t", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("\n")
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
}

//To check `race` condition run `$ go run -race exercise_60.go`
