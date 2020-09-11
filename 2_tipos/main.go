package main

import (
    "fmt"
    "reflect"
)

func main() {

    // len("string") tamanho da string
    // "string"[1] seleciona segundo caracter da string
    s := "string"

    // uint8 (byte), uint16, uint32, uint64 - inteiros sem sinal
    // int8, int16, int32 (rune), int64 - inteiros com sinal
    i := 10

    // float32, float64
    // complex64, complex128
    // NaN Not a Number para números que não podem ser representados (0/0 e infinito)
    f := 1.2

    // booleano com os operadores lógicos:
    // && (e) || (ou) ! (negação)
    b := true

    fmt.Println(reflect.TypeOf(s)) // string
    fmt.Println(reflect.TypeOf(i)) // int
    fmt.Println(reflect.TypeOf(f)) // float
    fmt.Println(reflect.TypeOf(b)) // bool

    fmt.Println("concatenacao" + "de string") //concatenacaode string
    fmt.Println(1 + 1.1) // 2.1
}
