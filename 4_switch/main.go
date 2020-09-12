package  main

import "fmt"

func main() {
    fmt.Print("Entre um numero: ")

    var input int
    fmt.Scanf("%d", &input)

    switch input {
        case 0: fmt.Println("Zero")
        case 1: fmt.Println("Um")
        case 2: fmt.Println("Dois")
        case 3: fmt.Println("Três")
        case 4: fmt.Println("Quatro")
        default: fmt.Println("????")
    }
}

