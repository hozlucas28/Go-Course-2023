/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Los punteros me permiten evitar que los valores sean copiados cuando
 *		   son enviados a una función, utilizando la memoria/referencia donde
 *		   dichos valores se encuentran almacenados.
-------------------------------------------------------------------------- */

package main

import "fmt"

func modifyValue(myStr *string) {
	fmt.Printf("%p\n", myStr)
	*myStr = "¡Hola desde la función!"
}

func main() {
	myStr := "¡Hola mundo de Go!"

	fmt.Printf("%p\n", &myStr)
	fmt.Println("Antes de la función:", myStr)

	modifyValue(&myStr)
	fmt.Println("Despues de la función:", myStr)
}
