package main

import (
	"fmt"
)

func main() {
	a := (10 == 20)
	b := (10 <= 20)
	c := (10 >= 20)
	d := (10 != 20)
	e := (10 < 20)
	f := (10 > 20)
	fmt.Println(a, b, c, d, e, f)
}
