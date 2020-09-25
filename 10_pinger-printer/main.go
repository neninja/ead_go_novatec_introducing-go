package main

import (
	"fmt"
	"time"
)

// sincronização de go routines através de canais
func main() {
	var c chan string = make(chan string)

	// inicia go routines
	go pinger(c)
	go printer(c)

	// segura a função main para finalizar as goroutines
	var input string
	fmt.Scanln(&input)
}

// a função só atribui ao canal
// caso não seja especificado, o canal é tido como bidirecional
// func pinger(c chan string) {
func pinger(c chan<- string) {
	// loop infinito
	for i := 0; ; i++ {
		// sempre que possível, atribuir o valor "ping" a c
		c <- "ping"
	}
}

// a função só remove do canal
// caso não seja especificado, o canal é tido como bidirecional
// func printer(c chan string) {
func printer(c <-chan string) {
	// loop infinito
	for {
		// sempre que possível, remover o valor
		// de c e atribuir à msg
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
