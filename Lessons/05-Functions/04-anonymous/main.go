/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Las funciones anónimas se ejecutan automáticamente si esta no se
 * 		   encuentra ligada a una variable, en caso de que si la misma debe
 *		   ser invocada a través de dicha variable como si se tratase de una
 *		   función de tipo flecha en JavaScript.
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	func() {
		fmt.Println("¡Hola desde una función anónima (ejecutada automáticamente)!")
	}()

	myFunction := func(name string) string {
		fmt.Printf("¡Hola %s desde una función anónima!\n", name)
		return name
	}

	myReturn := myFunction("Lucas")
	fmt.Println(myReturn)
}
