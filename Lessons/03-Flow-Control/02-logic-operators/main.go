package main

import "fmt"

func main() {
	// Negación
	fmt.Printf("• Negación:\n")
	fmt.Println(!true)
	fmt.Println(!false)
	fmt.Printf("\n")

	// Conjunción
	fmt.Printf("• Conjunción:\n")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Printf("\n")

	// Disyunción
	fmt.Printf("• Disyunción:\n")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Printf("\n")
}
