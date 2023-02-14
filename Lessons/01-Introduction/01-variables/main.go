/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Si no se le asigna ningún valor a las variables estas asumen un
 *		   valor por defecto, según el tipo de dato.
 *
 *
 * IMPORTANTE:
 *  			  - <:=> = Me almacenar en una variable cualquier tipo
 *						   de dato.
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	var firstName string = "Alex"
	var secondName string = "Alex"

	age := 20
	fmt.Println(firstName, secondName, age)

	var a int     // Valor por defecto <0>.
	var b float64 // Valor por defecto <0>.
	var c string  // Valor por defecto <"">.
	var d bool    // Valor por defecto <false>.
	fmt.Println(a, b, c, d)

	const pi = 3.141592
	fmt.Println(a, b, c, d)
}
