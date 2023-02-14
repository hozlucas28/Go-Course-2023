/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Los tipos de elementos del arreglo deben coincidir y se debe
 *		   establecer el tipo de dato de dichos elementos en la inicialización
 *		   del arreglo.
 *
 *
 * IMPORTANTE:
 *  			  - <len(<Array>)> = Devuelve el tamaño del arreglo o slice.
 *  			  - <<Array>[<START>:<END - 1>]> = Copia los elementos del
 * 												   arreglo según el rango
 *												   indicado.
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	// Inicialización
	var numbers [5]int
	fmt.Println(numbers)

	// Modificación de elementos
	numbers[0] = 10
	numbers[1] = 1
	numbers[2] = 3
	numbers[3] = 5
	numbers[4] = 7
	fmt.Println(numbers)

	// Referenciar elemento
	fmt.Println(numbers[3])

	// Inicialización alternativa con valores predefinidos
	names := [3]string{"Lucas", "Nahuel", "Alex"}
	fmt.Println(names)

	// Longitud dinámica
	colors := [...]string{"Rojo", "Verde", "Negro", "Azul"}
	fmt.Println(colors, len(colors))

	// Indicé definido
	coins := [...]string{
		0: "Dólar",
		2: "Peso",
		3: "Euro",
		5: "Libra",
	}
	fmt.Println(coins, len(coins))

	// Copia de elementos
	subCoins := coins[2:4]
	fmt.Println(subCoins, len(subCoins))
}
