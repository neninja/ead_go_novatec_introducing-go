package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person
type ByAge []Person

// sort tem a interface que exige Len, Less, Swap

// Como medir o tamanho
func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByAge) Len() int {
	return len(ps)
}

// Como saber quem é menor que quem
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

// Como trocar os valores
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids) // [{Jack 10} {Jill 9}]

	sort.Sort(ByAge(kids))
	fmt.Println(kids) // [{Jill 9} {Jack 10}]
}
