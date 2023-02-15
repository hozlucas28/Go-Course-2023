/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Para demostrar el funcionamiento de una función recursiva se
 *		   buscara obtener el factorial de un número.
-------------------------------------------------------------------------- */

package main

import "fmt"

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	f := number * factorial(number-1)

	return f
}

func main() {
	fmt.Println(factorial(5))
}
