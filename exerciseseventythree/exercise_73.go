package main

import (
	"fmt"

	"github.com/sirajummprince/go-exercises/exerciseseventythree/dog"
	//"./dog"
)

func ExampleDog() {
	dogYears := dog.Years(10)
	fmt.Println(dogYears)
	//Output:
	//70
}

func main() {
	dogYears := dog.Years(7)
	fmt.Println(dogYears)
}
