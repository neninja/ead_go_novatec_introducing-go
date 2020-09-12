package  main

import (
    "fmt"
    "strconv"
)

func main() {
    for i := 1; i < 100; i++ {
        var message string
        if i % 3 == 0 {
            message = "Fizz"
        }
        if i % 5 == 0 {
            message += "Buzz"
        }
        if message == "" {
            message = strconv.Itoa(i)
        }
        fmt.Println(message)
    }
}

