package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second): // a cada segundo esse case está pronto
				fmt.Println("timeout")
			default:
				time.Sleep(time.Second) // default a cada for caso nenhum case seja instantaneamente selecionado
				fmt.Println("nothing ready")
			}
		}
	}()

	// segura a função main para finalizar as goroutines
	var input string
	fmt.Scanln(&input)
}
