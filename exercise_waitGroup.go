// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// var wg sync.WaitGroup

// func foo() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("foo", i)
// 	}
// 	wg.Done()
// }

// func bar() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("bar", i)
// 	}
// }

// func main() {
// 	fmt.Println("OS:", runtime.GOOS)
// 	fmt.Println("ARCH:", runtime.GOARCH)
// 	fmt.Println("CPUs:", runtime.NumCPU())
// 	fmt.Println("Routines:", runtime.NumGoroutine())

// 	wg.Add(1)
// 	go foo()
// 	bar()

// 	fmt.Println("OS:", runtime.GOOS)
// 	fmt.Println("ARCH:", runtime.GOARCH)
// 	fmt.Println("CPUs:", runtime.NumCPU())
// 	fmt.Println("Routines:", runtime.NumGoroutine())
// }

// package main

// import "fmt"

// type person struct {
// 	first string
// }

// type human interface {
// 	speak()
// }

// func (p *person) speak() {
// 	fmt.Println("Hello")
// }

// func saySomething(h human) {
// 	h.speak()
// }

// func main() {
// 	p1 := person{
// 		first: "James",
// 	}

// 	saySomething(&p1)

// 	// does not work
// 	// saySomething(p1)

// 	p1.speak()
// }

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
	fmt.Println("Go Routines:\t", runtime.NumGoroutine())
	fmt.Println("\n")

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go Routines:\t", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("\n")
	fmt.Println("Counter:\t", counter)
	fmt.Println("Go Routines:\t", runtime.NumGoroutine())
}

//To check `race condition` run `$ go run -race exercise_58.go`
