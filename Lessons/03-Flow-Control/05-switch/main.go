/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   En caso del lenguaje Go no es necesario el break, ya que el
 *		   mismo no evalúa las demás condiciones al encontrar el primer
 *		   caso que cumpla dicha condición.
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	var vocal string

	fmt.Println("• Ingrese una letra:")
	fmt.Scanln(&vocal)

	// Primer método
	// switch {
	// case vocal == "a" || vocal == "A":
	// 	fmt.Printf("¡%s es vocal!", vocal)
	// case vocal == "e" || vocal == "E":
	// 	fmt.Printf("¡%s es vocal!", vocal)
	// case vocal == "i" || vocal == "I":
	// 	fmt.Printf("¡%s es vocal!", vocal)
	// case vocal == "o" || vocal == "O":
	// 	fmt.Printf("¡%s es vocal!", vocal)
	// case vocal == "u" || vocal == "U":
	// 	fmt.Printf("¡%s es vocal!", vocal)
	// default:
	// 	fmt.Printf("¡%s no es vocal!", vocal)
	// }

	// Segundo método
	switch vocal {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Printf("¡%s es vocal!", vocal)
	default:
		fmt.Printf("¡%s no es vocal!", vocal)
	}
}
