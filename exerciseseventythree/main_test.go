package main

import (
	"testing"

	"github.com/sirajummprince/go-exercises/exerciseseventythree/dog"
)

func TestDog(t *testing.T) {
	dy := dog.Years(10)
	if dy != 70 {
		t.Error("Something  went wrong. Expected", 70, "Got %v", dy)
	}
}

func BenchmarkDOg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.Years(10)
	}
}
