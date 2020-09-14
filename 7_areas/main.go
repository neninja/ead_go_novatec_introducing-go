package main

import (
	"fmt"
	"math"
)

// Structs são somente dados e suas as funções (métodos) que possuem
// (c *Circle) são seus métodos
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

type MultiShape struct {
	shapes []Shape
}

// Interfaces são especificações de "checklists" de métodos que devem ser
// implementados por structs
type Shape interface {
	area() float64
}

// antes de virar método de Circle
//func circleArea(c *Circle) float64 {
//	return math.Pi * c.r * c.r
//}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (m *MultiShape) area() (area float64) {
	for _, s := range m.shapes {
		area += s.area()
	}
	return
}

func main() {
	// c := new(Circle) // cria ponteiro

	// c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5} // na ordem de definição
	// c := &Circle{0, 0, 5} // criação do ponteiro (igual ao new())

	// antes de virar método de Circle
	// fmt.Println(circleArea(&c))
	fmt.Println(c.area())

	r := Rectangle{0, 10, 0, 10}
	fmt.Println(r.area())

	// duck typing:
	// se possui area() float64 -> é Shape
	fmt.Println(totalArea(&c, &r))

	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{0, 0, 5},
			&Rectangle{0, 10, 0, 10},
			// Circle{0, 0, 5}, // ¯\_(ツ)_/¯
			// Rectangle{0, 10, 0, 10},
		},
	}

	fmt.Println(multiShape.area())
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func totalArea(shapes ...Shape) (area float64) {
	for _, s := range shapes {
		area += s.area()
	}
	return
}
