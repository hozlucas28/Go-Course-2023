/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Los "maps" son elementos desordenados en donde para acceder
 *		   a un valor se debe referenciar su clave ligada. A los mapas no
 *		   se le puede definir una capacidad. Si buscamos obtener el valor
 *		   asociado a una llave inexistente el retorno sera un <String> vacío,
 *		   con existencia <false>.
 *
 *
 * IMPORTANTE:
 *  			  - El indicé de los "mapas" comienza en 1.
 *  			  - El segundo valor de retorno de un <map> al obtener un valor
 *					es un booleano que indica si dicho valor existe.
 *  			  - <delete(<MAP>, <KEY>)> = Me permite eliminar el valor
 *											 de la llave ligada (argumento)
 *											 de dicho mapa (argumento).
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	// Inicialización
	days := make(map[int]string) // make(map[<DATA TYPE OF KEYS>]<DATA TYPE OF VALUES>)
	fmt.Println(days)

	// Agregar valores
	days[1] = "Lunes" // <MAP>[<KEY>] = <VALUE>
	days[2] = "Martes"
	days[3] = "Miércoles"
	days[4] = "Jueves"
	days[5] = "Viernes"
	days[6] = "Sabado"
	days[7] = "Domingo"
	fmt.Println(days)

	// Obtener valor
	fmt.Println(days[4])

	// Eliminar valor
	delete(days, 2)
	fmt.Println(days, len(days))

	// Mapa con arreglos
	students := make(map[string][]int)
	students["Lucas"] = []int{20, 15, 16}
	students["Nahuel"] = []int{32, 11, 21}
	fmt.Println(students, len(students))
	fmt.Println(students["Lucas"][0])
}
