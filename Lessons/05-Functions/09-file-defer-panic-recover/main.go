/* --------------------------------------------------------------------------
 * IMPORTANTE:
 *  			  - <panic()> = Realiza una detención brusca del
 *								programa e informa la posible ubicación
 *								del error, como si se trace de una <exception>
 *								en Java.
 *  			  - <recover()> = Me permite controlar los "panics" de mi
 *								  programa, evitando que se detenga bruscamente.
 *  			  - <Open(<FILE>)> = Abre el archivo (argumento).
 *  			  - <defer <FUNCTION>> = Ejecuta dicha función (argumento)
 *										 una vez que la función padre ha terminado
 *										 de ejecutarse.
 *  			  - <<FILE>.Read(<ARRAY>)> = Lee el contenido del archivo (argumento).
-------------------------------------------------------------------------- */

package main

import (
	"fmt"
	"os"
)

func main() {
	// Controlar "panics"
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Al parecer el programa no ha finalizado de forma correcta")
		}
	}()

	// Abrir archivo
	if file, error := os.Open("hell.txt"); error != nil {
		panic("¡Error! No se ha podido abrir el archivo: hello.txt")
	} else {
		// Cerrar archivo
		defer func() {
			fmt.Println("¡Cerrando archivo!")
			file.Close()
		}()

		// Leer contenido del archivo
		content := make([]byte, 254)
		lenght, _ := file.Read(content)
		fileContent := string(content[:lenght])
		fmt.Println(fileContent)
	}
}
