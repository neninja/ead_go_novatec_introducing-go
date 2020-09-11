package main

import "fmt"

func main() {
	fmt.Print("Entre.com a temperatura em °C: ")

	var celsius float64
	fmt.Scanf("%f", &celsius) // igual a c ಥ_ಥ
	fahrenheit := (celsius / 5 * 9) + 32

	fmt.Printf("%.2f°C -> %.2f°F\n", celsius, fahrenheit)
}
