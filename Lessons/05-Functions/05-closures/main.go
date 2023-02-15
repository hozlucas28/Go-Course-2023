/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Las funciones <Closure> retornan una función anónima creada
 *		   dinámicamente.
 *
 *
 * IMPORTANTE:
 *  			  - <Repeat(<STRING>, <N>)> = Repita la cadena N (argumento) veces.
-------------------------------------------------------------------------- */

package main

import (
	"fmt"
	"strings"
)

// Función de tipo <Closure>
func repeat(number int) func(myString string) string {
	return func(myString string) string { // Función anónima retornada.
		return strings.Repeat(myString, number)
	}
}

func main() {
	var numberOfRepeats int = 3

	repeatReturn := repeat(numberOfRepeats)
	fmt.Println(repeatReturn("Lucas"))
}
