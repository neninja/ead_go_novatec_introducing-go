package math

//func main() {
//	xs := []float64{98, 93, 77, 82, 83}
//	fmt.Println(Average(xs))
//}

// Average Calcula média de uma série de números
func Average(xs []float64) (total float64) {
	total = 0.0

	for _, nota := range xs {
		total += nota
	}

	total /= float64(len(xs))
	// parametro de retorno nomeado
	return
}

//func Average(xs []float64) float64 {
//	total := 0.0
//
//	for _, nota := range xs {
//		total += nota
//	}
//
//	return total / float64(len(xs))
//}
