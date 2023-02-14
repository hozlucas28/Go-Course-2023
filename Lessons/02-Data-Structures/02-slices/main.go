/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Los "slices" son mutables, es decir, se le pueden agregar
 * 		   elementos ampliando el tamaño del mismo. Hay que tener en cuenta
 *		   que las modificaciones aplicadas a un "slice" hijo también son aplicadas
 *		   al "slice" padre.
 *
 *
 * IMPORTANTE:
 *  			  - <append(<SLICE>, <ELEMENT>)> = Devuelve un nuevo "slice",
 *												   es decir, aquel "slice" (argumento)
 *												   con el elemento (argumento) añadido
 *												   al final.
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	// Inicialización
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

	// Agregar elemento
	numbers = append(numbers, 4)
	numbers = append(numbers, 5)
	fmt.Println(numbers, len(numbers))

	// Extraer elementos
	subNumbers := numbers[:2]
	subNumbers[0] = 100
	fmt.Println(subNumbers)
	fmt.Println(numbers)

	// Mostrar capacidad y memoria
	months := []string{"Enero", "Febrero", "Marzo"}
	fmt.Printf("Longitud: %v | Capacidad: %v | Memoria: %p\n", len(months), cap(months), months)

	months = append(months, "Abril")
	fmt.Printf("Longitud: %v | Capacidad: %v | Memoria: %p\n", len(months), cap(months), months)
}
