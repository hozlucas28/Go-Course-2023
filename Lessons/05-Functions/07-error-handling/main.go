package main

import (
	"errors"
	"fmt"
)

func division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("¡Error! No se puede dividir por cero.")
	}

	return a / b, nil
}

func main() {
	result, error := division(4, 0)

	if error == nil {
		fmt.Printf("• El resultado es: %d\n", result)
	} else {
		fmt.Println(error)
	}
}
