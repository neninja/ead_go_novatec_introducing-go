package  main

import "fmt"

func main() {
    var x [5]int // array contendo 5 posições ocm inteiros
    // Arrays obrigatoriamente tem tamanho fixo, diferente de slices

    fmt.Println(len(x)) // tamanho do array (5)

    x[4] = 100 // 100 na quinta posição (inicia na 0)

    fmt.Println(x) // [0 0 0 0 100]

    // funciona semelhante ao foreach
    for i, value := range x {
        fmt.Println(i, value)
        // 0 0
        // 1 0
        // 2 0
        // 3 0
        // 4 100
    }

    // array com inicialização
    y := [5]int { 12, 54, 21, 4, 7 }

    // caso não seja usado i
    // "i declared but not used"
    for _, value := range y {
        fmt.Println(value)
        // 12
        // 54
        // 21
        // 4
        // 7
    }
}

