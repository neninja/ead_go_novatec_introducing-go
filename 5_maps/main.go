package main

import "fmt"

func main() {
	// dicionários/arrays associativos/chave valor
	x := make(map[string]int)
	x["key"] = 5
	x["yek"] = 5
	fmt.Println(x) // map[key:5 yek:5]
	fmt.Println(x["key"])
	delete(x, "yek")
	fmt.Println(x) // map[key:5]

	valor, ok := x["key"]
	fmt.Println(valor, ok)            // 5 true
	valorInexistente, ok := x["key2"] // ok testa se a chave existe
	fmt.Println(valorInexistente, ok) // 0 false
	x["key3"] = 0
	valorZerado, ok := x["key3"] // ok testa se a chave existe
	fmt.Println(valorZerado, ok) // 0 true

	if name, ok := x["inexistente"]; ok {
		fmt.Println(name, ok) // não entra nesse if
	}

	meses := map[string]int{
		"janeiro":   1,
		"fevereito": 2,
		"marco":     3,
		"abril":     4,
		"maio":      5,
		"junho":     6,
		"julho":     7,
		"agosto":    8,
		"setembro":  9,
		"outubro":   10,
		"novembro":  11,
		"dezembro":  12,
	}

	fmt.Println(meses)

}
