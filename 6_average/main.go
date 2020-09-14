package main

import "fmt"

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
}

// parametro de retorno nomeado
func average(xs []float64) (total float64) {
	total = 0.0

	for _, nota := range xs {
		total += nota
	}

	total /= float64(len(xs))
	return
}

//func average(xs []float64) float64 {
//	total := 0.0
//
//	for _, nota := range xs {
//		total += nota
//	}
//
//	return total / float64(len(xs))
//}
