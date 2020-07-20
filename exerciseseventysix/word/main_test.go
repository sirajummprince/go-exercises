package word

import (
	"testing"
)

func TestCount(t *testing.T) {
	c := Count("Sirajum Munir Prince")
	if c != 3 {
		t.Error("Wrong count. Expected", 3, "Got", c)
	}
}

func TestUseCount(t *testing.T) {
	c := UseCount("Sirajum Munir Prince Prince")
	for i, v := range c {
		switch i {
		case "Sirajum", "Munir":
			if v != 1 {
				t.Error("Wrong use count.Expected", 1, "Got", c)
			}
		case "Prince":
			if v != 2 {
				t.Error("Wrong use count.Expected", 2, "Got", c)
			}
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(Count.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(UseCount.SunAlso)
	}
}

func ExampleCount() {
	Count("Sirajum Munir Prince Prince")
	//Output:
	//4
}

func ExampleUseCount() {
	UseCount("Sirajum Munir Prince Prince")
	//Output:
	//1 Sirajum
	//1 Munir
	//2 Prince
}
