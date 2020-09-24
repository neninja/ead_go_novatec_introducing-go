package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Tests", "ests"))  // true
	fmt.Println(strings.Contains("Testes", "test")) // false

	fmt.Println(strings.Count("testes", "t")) // 2
	fmt.Println(strings.Count("Testes", "t")) // 1

	fmt.Println(strings.HasPrefix("Testes", "tes")) // false
	fmt.Println(strings.HasPrefix("Testes", "Tes")) // true

	fmt.Println(strings.HasSuffix("Testes", "tes")) // true
	fmt.Println(strings.HasSuffix("Testes", "Tes")) // false

	fmt.Println(strings.Index("Testes", "e")) // 1
	fmt.Println(strings.Index("Testes", "s")) // 2

	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-")) // a-b-c

	fmt.Println(strings.Repeat("qq", 5)) // qqqqqqqqqq

	fmt.Println(strings.Replace("qqqqqqqqqq", "qqq", "bbb", 2)) // bbbbbbqqqq

	fmt.Println(strings.Split("a-b-c", "-")) // [ a b c ]

	fmt.Println(strings.ToLower("TEStes")) // testes
	fmt.Println(strings.ToUpper("TEStes")) // TESTES
}
