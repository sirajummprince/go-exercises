package mymath

import (
	"fmt"
	"testing"
)

//TestCenteredAvg tests if the CenteredAvg works properly
func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		result float64
	}

	tests := []test{
		test{
			[]int{1, 3, 5, 7, 9},
			5,
		},
		test{
			[]int{2, 4, 6, 8, 10},
			6,
		},
		test{
			[]int{11, 13, 15, 17, 19},
			15,
		},
		test{
			[]int{20, 22, 24, 26, 28, 30},
			25,
		},
	}

	for _, v := range tests {
		r := CenteredAvg(v.data)
		if r != v.result {
			t.Error("Expected", v.result, "Got", r)
		}
	}
}

//BenchmarkCenteredAvg benchmarks the CenteredAvg
func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{2, 3, 4, 5, 6, 7}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}

//ExmapleCenteredAvg displays an example regarding CenteredAvg
func ExampleCenteredAvg() {
	xi := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(CenteredAvg(xi))
	//Output:
	//4.5
}
