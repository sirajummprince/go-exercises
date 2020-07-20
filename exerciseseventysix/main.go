package main

import (
	"fmt"

	"github.com/sirajummprince/go-exercises/exerciseseventysix/quote"
	"github.com/sirajummprince/go-exercises/exerciseseventysix/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
