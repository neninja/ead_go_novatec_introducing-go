package main

import (
	"errors"
	"fmt"
)

// http://goporexemplo.golangbr.org/errors.html
// https://golang.org/doc/effective_go.html#errors

func divide(dividendo float64, divisor float64) (float64, error) {
	if divisor == 0 {
		return -1, errors.New("não pode trabalhar com 0")
	}
	return divisor / dividendo, nil
}

func main() {
	// forma 1
	resultado1, err := divide(15, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado1)
	}

	// forma 2
	if resultado2, err := divide(15, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado2)
	}
}
