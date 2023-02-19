package main

import (
	"fmt"
	"math"
)

/* -------------------------------- Interfaz -------------------------------- */

type IFigure interface {
	getArea() float32
	getPerimeter() float32
}

func figureLenghts(iF IFigure) {
	fmt.Println(iF)
	fmt.Println(iF.getArea())
	fmt.Println(iF.getPerimeter())
}

/* --------------------------------- Objetos -------------------------------- */

type Square struct {
	side float32
}

func (s *Square) getArea() float32 {
	return s.side * 2
}

func (s *Square) getPerimeter() float32 {
	return s.side * 4
}

type Circle struct {
	radio float32
}

func (c *Circle) getArea() float32 {
	return math.Pi * (c.radio * c.radio)
}

func (c *Circle) getPerimeter() float32 {
	return 2 * math.Pi * c.radio
}

/* --------------------------- Programa Principal --------------------------- */

func main() {
	square := Square{side: 4}
	circle := Circle{radio: 5}

	figureLenghts(&square)
	figureLenghts(&circle)
}
