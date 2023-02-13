package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	var a, b int

	fmt.Println("• Ingrese el primer número:")
	fmt.Scanln(&a)

	fmt.Println("• Ingrese el segundo número:")
	fmt.Scanln(&b)

	var result int = add(a, b)
	fmt.Printf("EL RESULTADO ES %d.", result)
}
