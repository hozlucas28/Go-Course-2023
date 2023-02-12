package main

import "fmt"

func greet(name string) {
	fmt.Printf("¡Hola %s!\n", name)
}

func add(a, b int) int {
	return a + b
}

func main() {
	greet("Lucas")
	greet("Nahuel")

	result := add(10, 20)
	fmt.Println("• La suma es:", result)
}
