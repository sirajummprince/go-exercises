package main

import (
	"fmt"
)

func foo() {
	fmt.Println("This is foo")
}

func bar() {
	fmt.Println("This is bar")
}

func main() {
	fmt.Println("Before defer:")
	foo()
	bar()
	fmt.Println("After defer:")
	defer foo()
	bar()
}
