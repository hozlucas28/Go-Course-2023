/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Las interfaces me permiten manejar los métodos idénticos de las
 *		   estructuras.
-------------------------------------------------------------------------- */

package main

import "fmt"

// Interface
type IAnimal interface {
	move() string
}

// Método de la interfaz
func moveAnimal(animal IAnimal) {
	fmt.Println(animal.move())
}

type Dog struct{}
type Fish struct{}
type Bird struct{}

func (d *Dog) move() string {
	return "Soy un perro y camino."
}

func (f *Fish) move() string {
	return "Soy un pez y nado."
}

func (b *Bird) move() string {
	return "Soy un pájaro y vuelo."
}

func main() {
	dog := Dog{}
	moveAnimal(&dog)

	fish := Fish{}
	moveAnimal(&fish)

	bird := Bird{}
	moveAnimal(&bird)
}
