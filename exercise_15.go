package main

import (
	"fmt"
)

func main() {
	byear := 1993
	for {
		if byear > 2020 {
			break
		}
		fmt.Println(byear)
		byear++
	}

}
