package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	// lista de sites
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
		"http://www.apple.com",
	}

	// cria canal do tipo da struct, obrigando inicialização da mesma
	results := make(chan HomePageSize)

	for _, url := range urls {
		// função anônima como goroutine, necessitando da variavel url, sendo
		// passada como parâmetro
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			bs, err := ioutil.ReadAll(res.Body) // ¯\_(ツ)_/¯
			if err != nil {
				panic(err)
			}
			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize
	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("The biggest home page:", biggest.URL)
}
