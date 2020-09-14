package main

import "fmt"

// a função anonima criada dentro de uma var
// utiliza outra var interna
// "A função juntamente com x formam a closure"
func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}
