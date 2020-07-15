package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int = 10
var wg sync.WaitGroup

func foo() {
	for i := 0; i < count; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < count; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}

func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())
}
