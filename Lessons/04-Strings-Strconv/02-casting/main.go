/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Para demostrar el casteo de datos se buscara resolver una
 *		   operación de suma en base a la entrada del usuario.
 *
 *
 * IMPORTANTE:
 *  			  - <Atoi(<STRING>)> = Castea el <String> dado a <Integer>.
-------------------------------------------------------------------------- */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(operation string) int {
	var result int
	operationArray := make([]string, 0)

	operationArray = strings.Split(operation, "+")

	for _, v := range operationArray {
		number, error := strconv.Atoi(v)

		// Manejo de errores
		if error == nil {
			result += number
		} else {
			result = -1
			// fmt.Println(error)
			fmt.Println("¡Error! Debe ingresar números enteros y operaciones de suma (+).")
			break
		}
	}

	return result
}

func main() {
	var operation string

	fmt.Println("• Ingrese la operación:")
	fmt.Scanln(&operation)

	result := add(operation)
	if result != -1 {
		fmt.Printf("¡El resultado de la operación es %d!", result)
	}
}
