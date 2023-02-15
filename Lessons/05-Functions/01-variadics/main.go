/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Las funciones variádicas son aquellas funciones que reciben un
 *		   número indefinido de parámetros.
-------------------------------------------------------------------------- */

package main

import "fmt"

func add(numbers ...int) int {
	var result int

	for _, v := range numbers {
		result += v
	}

	return result
}

func main() {
	result := add(1, 3, 4, 7, 10, 21)
	fmt.Println("• Resultado de la suma:", result)
}
