package main

import (
	"fmt"
)

type customErr struct {
	err string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("then received the error %v", ce.err)
}

func foo(e error) {
	fmt.Println("foo ran first", e)
}

func main() {
	c1 := customErr{
		err: `finally printed the error`,
	}
	foo(c1)
}
