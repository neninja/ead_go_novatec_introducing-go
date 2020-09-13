package  main

import "fmt"

func main() {
    // slice de inteiros contendo inicialmente 5 posições
    x := make([]int, 5)

    fmt.Println(len(x)) // tamanho atual do slice (5)

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

    y := append(x, 59) // novo array com base no primeiro
    fmt.Println(x) // [0 0 0 0 100] não é afetado com o append

    // caso não seja usado i
    // "i declared but not used"
    for _, value := range y {
        fmt.Println(value)
        // 0
        // 0
        // 0
        // 0
        // 100
        // 59
    }

    // ao copiar diretamente, é passado a referência do original
    // alterar em um afeta o outro
    z := x
    z[4] = 99
    fmt.Println(x) // [0 0 0 0 99] é afetado com a operação
    fmt.Println(z) // [0 0 0 0 99] slice "novo" original

    // e necessário usar copy() para evitar isso
    zz := make([]int, 5)
    copy(zz, x) // copy(destino, fonte)
    z[4] = 11
    fmt.Println(x) // [0 0 0 0 99] não é afetado com a operação
    fmt.Println(zz) // [0 0 0 0 11] slice novo
}

