package main

import (
	"fmt"
	"testing"

	//"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/02/01-code-starting/quote"
	//"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/02/01-code-starting/word"
	"./quote"
	"./word"
)

func Test(t *testing.T) {

}

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
