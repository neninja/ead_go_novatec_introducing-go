package main

import "fmt"

func main() {
    fmt.Print("Entre um numero: ")

	var input float64
    fmt.Scanf("%f", &input) // igual a c ಥ_ಥ
    output := input * 2

    fmt.Println(output)
}
