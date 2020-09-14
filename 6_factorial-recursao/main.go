package main

import "fmt"

func main() {
	// 4 * 3 * 2 * 1 = 24
	fmt.Println(factorial(4))
}

// substitui o for
func factorial(x uint) uint {
	// quebra o loop infinito
	// deixando de retornar uma nova factorial(x)
	if x == 0 {
		return 1
	}

	// invoca ela mesma, com um "parâmetro novo"
	return x * factorial(x-1)
}
