/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Al trabajar con <forEach> debo definir dos variables, la primera
 *		   sera el indicé y la segunda el elemento del arreglo a recorrer.
 *
 *
 * IMPORTANTE:
 *  			  - Al colocar <_> en la variable del indicé del <forEach>
 *					esta no lanzara un warning por no utilizarla.
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	names := [...]string{"Lucas", "Nahuel", "Juan"}

	// Ejemplo de tipo <for>
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// Ejemplo de tipo <forEach>
	for i, element := range names {
		fmt.Println(i, element)
	}

	// Ejemplo de tipo <forEach>, sin tomar en cuenta el indicé
	for _, element := range names {
		fmt.Println(element)
	}
}
