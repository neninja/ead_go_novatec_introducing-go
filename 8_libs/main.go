package main

// para ver a documentação local use:
// godoc -http=:6060
// http://localhost:6060/pkg/github.com/nenitf/ead_go_novatec_introducing-go/6_average/

// Caso não possua godoc, instale com:
// go get -v golang.org/x/tools/cmd/godoc

import (
	"fmt"

	math "github.com/nenitf/ead_go_novatec_introducing-go/6_average"
)

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(math.Average(xs))
}
