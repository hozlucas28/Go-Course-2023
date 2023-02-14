/* --------------------------------------------------------------------------
 * IMPORTANTE:
 *  			  - <%t> = Cuando el tipo de dato es <Bool>.
 *  			  - <%s> = Cuando el tipo de dato es <String>.
 *  			  - <%d> = Cuando el tipo de dato es <Integer>.
 *  			  - <%v> = Cuando el tipo de dato no se conoce.
 *  			  - <%p> = Formatea el mensaje con la dirección de
 *						   memoria de la variable ligada.
 *  			  - <%T> = Formatea el mensaje con el tipo de dato
 * 						   de la variable.
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	hello := "Hola"
	world := "Mundo"

	// Mensajes
	fmt.Print(hello)
	fmt.Print(world)

	// Mensajes por línea
	fmt.Println(hello)
	fmt.Println(world)

	// Mensajes formateados
	name := "Lucas"
	age := 20
	fmt.Printf("Hola %s, tu edad es %d años.\n", name, age)
	fmt.Printf("Hola %v, tu edad es %v años.\n", name, age)

	// Retorno de mensaje formateado
	message := fmt.Sprintf("Hola %s, tu edad es %d años.", name, age)
	fmt.Println(message)

	fmt.Printf("Hola %T, tu edad es %T años.\n", name, age)

	// Ingreso de datos por teclado
	fmt.Print("• Ingrese otro nombre:")
	fmt.Scanln(&name)

	fmt.Printf("El nombre ingresado es: %v", name)
}
