package main

import "fmt"

func add(message string, numbers ...int) (string, int) {
	var result int

	for _, v := range numbers {
		result += v
	}

	fmt.Printf("• Resultado de la %s es: %d\n", message, result)
	return message, result
}

func main() {
	message, result := add("suma", 1, 3, 4, 7, 10, 21)
	fmt.Println(message, result)
}
