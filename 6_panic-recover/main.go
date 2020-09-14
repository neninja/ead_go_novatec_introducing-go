package main

import "fmt"

func main() {
	// primeiroExemplo()
	segundoExemplo()
}

func primeiroExemplo() {
	panic("DIE")
	str := recover() // isto jamais ocorrerá
	fmt.Println(str)
}

func segundoExemplo() {
	defer func() {
		str := recover() // tratamento do panic
		fmt.Printf("imprissão do panic: %s", str)
	}()
	panic("DIE")
}
