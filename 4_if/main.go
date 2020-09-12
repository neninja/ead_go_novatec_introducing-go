package  main

import "fmt"

func main() {
    fmt.Print("Entre um numero: ")

	var input int
    fmt.Scanf("%d", &input)

    if input % 2 == 0 {
        fmt.Println("even") // par
    } else {
        fmt.Println("odd") // impar
    }
}

