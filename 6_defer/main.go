package main

import "fmt"

func main() {
	/* escrever dessa forma */
	defer second()
	first()

	/* é igual a */
	// first()
	// second()
	// defer joga a instrução para o final da função
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}
