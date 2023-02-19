/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   En el lenguaje Go no existe el bucle <while>.
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	// Bucle infinito
	/*
		for {
			println("¡Hola Mundo!")
		}

		Bucle tipo <while>
		numbers := 12455
		counter := 0

		for numbers > 0 {
			numbers /= 10
			counter++
		}
		fmt.Println(counter)
	*/

	// Bucle tipo <for>
	for i := 0; i < 32; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
