package main

import "fmt"

// variavel *int -> ponteiro inteiro
// *variavel = 1 -> atribuir valor para a variavel do ponteiro (desreferencia)
//  variavel = 1 -> NÃO funciona

// &x -> passa referencia da memoria da variável

func main() {
	x := 5
	one(&x)        // passa a referencia da variável (como ponteiro)
	fmt.Println(x) // 1

	y := new(int)   // cria *int
	one(y)          // y já é ponteiro
	fmt.Println(y)  // endereço de memoria
	fmt.Println(*y) // 1
}

// uma função que espera *int deve receber &variavel

func one(xPtr *int) {
	*xPtr = 1
}
