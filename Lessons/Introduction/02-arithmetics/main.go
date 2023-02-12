package main

import "fmt"

func main() {
	a := 20
	b := 10

	// Suma
	result := a + b
	fmt.Println("• Suma:", result)

	// Resta
	result = a - b
	fmt.Println("• Resta:", result)

	// Multiplicación
	result = a * b
	fmt.Println("• Multiplicación:", result)

	// División
	result = a / b
	fmt.Println("• División N°1:", result)

	// División
	var div float64 = 3.0 / 2.0
	fmt.Println("• División N°2:", div)

	// Módulo
	result = 3 % 2
	fmt.Println("• Módulo:", result)
}
