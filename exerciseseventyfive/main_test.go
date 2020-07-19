package main

import (
	"fmt"
	"testing"

	"github.com/sirajummprince/go-exercises/exerciseseventythree/dog"
)

func BenchmarkDog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.Years(10)
	}
}

//ExampleDog shows how dog package works
func ExampleDog() {
	fmt.Println(dog.Years(10))
	//Output:
	//70
}

func TestDog(t *testing.T) {
	dy := dog.Years(10)
	if dy != 70 {
		t.Error("Expected", 70, "Got ", dy)
	}
}
