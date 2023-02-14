package main

import "fmt"

func main() {
	var myReturn bool

	a := 4
	b := 5

	// Igualdad
	myReturn = a == b
	fmt.Printf("¿%d es igual a %d? %t\n", a, b, myReturn)

	// Distinto
	myReturn = a != b
	fmt.Printf("¿%d es distinto a %d? %t\n", a, b, myReturn)

	// Mayor que
	myReturn = a > b
	fmt.Printf("¿%d es mayor a %d? %t\n", a, b, myReturn)

	// Mayor o igual que
	myReturn = a >= b
	fmt.Printf("¿%d es mayor o igual a %d? %t\n", a, b, myReturn)

	// Menor que
	myReturn = a < b
	fmt.Printf("¿%d es menor a %d? %t\n", a, b, myReturn)

	// Menor o igual que
	myReturn = a <= b
	fmt.Printf("¿%d es menor o igual a %d? %t\n", a, b, myReturn)
}
