package math

import "fmt"

func ExampleAverage() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(Average(xs))

	xs = []float64{9, 13, 55, 94, 20}
	fmt.Println(Average(xs))
	// Output:
	// 86.6
	// 38.2
}
