package main

import "fmt"

func getQuotient(a, b int) int {
	return a / b
}

func getRemainder(a, b int) int {
	return a % b
}

func main() {
	var a, b int

	fmt.Println("• Ingrese el primer número:")
	fmt.Scanln(&a)

	fmt.Println("• Ingrese el segundo número:")
	fmt.Scanln(&b)

	var quotient int = getQuotient(a, b)
	var remainder int = getRemainder(a, b)
	fmt.Printf("EL COCIENTE ES %d Y EL RESTO ES %d.", quotient, remainder)

}
