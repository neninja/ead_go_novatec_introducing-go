package math

import (
	"fmt"
	"testing"
)

// table driven test
// https://github.com/golang/go/wiki/TableDrivenTests
func TestAverage(t *testing.T) {
	var tests = []struct {
		in  []float64
		out float64
	}{
		{[]float64{1, 2}, 1.5},
		{[]float64{1, 1, 1, 1, 1, 1, 1}, 1},
		{[]float64{-1, 1}, 0},
	}
	for _, test := range tests {
		got := Average(test.in)
		if got != test.out {
			t.Error("For", test.in, "want", test.out, "got", got)
		}
	}
}

/*
func TestAverage(t *testing.T) {
	got := Average([]float64{1, 2})
	want := 1.5
	if got != want {
		t.Error("Got", got, "want", want)
	}
}
*/

func ExampleAverage() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(Average(xs))

	xs = []float64{9, 13, 55, 94, 20}
	fmt.Println(Average(xs))
	// Output:
	// 86.6
	// 38.2
}
