package main

import "fmt"

func criaProximoNumeroPar() func() uint {
	i := uint(0)
	return func() (ret uint) {
		// como essa função sempre é chamada
		// o valor de i é persistido
		ret = i
		i += 2
		return
	}
}

func main() {
	// inicializa a função com i = 0
	proximoPar := criaProximoNumeroPar()

	// invoca somente a função anonima dentro de criaProximoNumeroPar()
	// não alterando i
	fmt.Println(proximoPar()) // 0
	fmt.Println(proximoPar()) // 2
	fmt.Println(proximoPar()) // 4
	fmt.Println(proximoPar()) // 6
}
