package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Error message")
	fmt.Println(err)
}
