package main

import (
	"fmt"
)

func main() {
	favSport := "cricket"
	switch favSport {
	case "football":
		fmt.Println("football is not my favourite")
	case "hockey":
		fmt.Println("hockey is not my favourite")
	case "cricket":
		fmt.Println("cricket is my favourite")
	}
}
