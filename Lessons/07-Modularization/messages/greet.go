/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Las constantes, variables, funciones, etc. que empiezan con
 *		   minúscula son privadas y solo se pueden reutilizar dentro del
 *		   mismo archivo/paquete, en este caso "messages".
-------------------------------------------------------------------------- */

package messages

import "fmt"

func Greet() { // Función pública.
	fmt.Println("¡Hola desde el paquete Mensajes!")
}

const message = "¡Hola desde mi constante!" // Constante privada.

func Print() { // Función pública.
	fmt.Println(message)
	fPrivate()
}

func fPrivate() { // Función privada
	fmt.Println("¡Hola desde mi función privada!")
}
