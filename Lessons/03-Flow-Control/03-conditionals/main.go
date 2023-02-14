/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Explicación realizada con el ejemplo de un restaurante con los
 *		   siguientes datos:
 *							 - Descuento del 10% hasta $100 de consumo.
 *							 - Descuento del 20% hasta $200 de consumo.
 *							 - Descuento del 30% mayor a $200 de consumo.
 *							 - IGV del 19%, respecto del consumo menos descuento.
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	var consumption, discount float32
	var discountStr string

	// Entrada
	fmt.Println("• Ingrese el consumo:")
	fmt.Scanln(&consumption)

	// Lógica
	if consumption > 0 {
		if consumption <= 100 {
			discount = consumption * 0.1
			discountStr = "10%"
		} else if consumption > 100 && consumption <= 200 {
			discount = consumption * 0.2
			discountStr = "20%"
		} else if consumption > 200 {
			discount = consumption * 0.3
			discountStr = "30%"
		}
	} else {
		fmt.Println("¡Error! El consumo debe ser mayor a $0")
		return
	}

	// Operaciones
	finalPrice := consumption - discount
	igv := finalPrice * 0.19
	finalPrice = finalPrice + igv

	// Salidas
	fmt.Println("================================")
	fmt.Printf("• IGV: $%0.2f\n", igv)
	fmt.Printf("• Consumo: $%0.2f\n", consumption)
	fmt.Printf("• Descuento: %s\n", discountStr)
	fmt.Printf("• Monto del Descuento: $%0.2f\n", discount)
	fmt.Printf("• Precio Final: $%0.2f\n", finalPrice)
	fmt.Println("================================")
}
