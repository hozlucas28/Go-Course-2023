/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   La función <make> se utiliza para definir un <Map> ó un <Slice>.
 *		   En el caso de un <Slice> esta me permite predefinir la longitud y
 *		   capacidad.
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	numbers := make([]int, 3, 3) // make(<MAP/SLICE>, <LENGTH>, <CAPACITY>)

	numbers[0] = 100
	numbers[1] = 200
	numbers[2] = 300

	numbers = append(numbers, 400)

	fmt.Println(numbers, len(numbers), cap(numbers))
}
