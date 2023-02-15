/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Las variables globales se define en el scope más alto del paquete
 *		   y pueden modificarse desde cualquier punto del mismo.
-------------------------------------------------------------------------- */

package main

import "fmt"

// Variable global
var message string

func fn1() {
	message = "¡Hola desde la función N°1!"
	fmt.Println(message)
}

func fn2() {
	message = "¡Hola desde la función N°2!"
	fmt.Println(message)
}

func main() {
	message = "¡Hola desde la función main!"
	fmt.Println(message)

	fn1()
	fn2()

	fmt.Println(message)
}
