package main

import "fmt"

func soma(numeros ...int) (total int) {
	for _, numero := range numeros {
		total += numero
	}
	return
}

func main() {
	fmt.Println(soma(1, 2, 3, 4))       // 10
	fmt.Println(soma(1, 2, 3, 4, 5, 6)) // 21
}
