package main

import (
	"fmt"
)

type customErr struct {
	err string
}

func (ce customErr) Error() string {
	return fmt.Sprintf(ce.err)
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	c1 := customErr{
		err: `Error occured in customErr`,
	}
	foo(c1)
}
