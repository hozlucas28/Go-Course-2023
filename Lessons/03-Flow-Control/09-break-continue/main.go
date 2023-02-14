package main

/* --------------------------------------------------------------------------
 * IMPORTANTE:
 *  			  - <Break> = Termina el bucle.
 *  			  - <Continue> = Salta a la siguiente iteración.
-------------------------------------------------------------------------- */

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			fmt.Println("Salto de iteración por <continue>")
			continue
		}

		if i == 8 {
			fmt.Println("Bucle terminado por <break>")
			break
		}

		fmt.Println(i)
	}
}
