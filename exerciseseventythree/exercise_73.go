package main

import (
	"fmt"

	"./dog"
)

func main() {
	dogYears := dog.Years(7)
	fmt.Println(dogYears)
}
