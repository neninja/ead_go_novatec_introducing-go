package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main conta como a go routine principal
// caso ela encerre todas outras encerram junto
func main() {
	// cria 10 go routines
	for i := 0; i < 10; i++ {
		go f(i)
	}

	// segura a função main para finalizar as goroutines
	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
