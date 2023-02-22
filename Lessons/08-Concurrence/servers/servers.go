/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   La concurrencia me permite ejecutar multiples instrucciones en
 *		   diferentes hilos para acortar tiempos, pero en caso de retornos
 *		   deberé implementar canales de comunicación para captar dichos
 *		   retornos.
 *
 *
 * IMPORTANTE:
 *  			  - <go> = Define la instrucción siguiente como una concurrencia.
 *  			  - <make(chan <DATA TYPE>)> = Me permite definir un canal que
 *											   establece la comunicación con las
 *											   concurrencias obteniendo y almacenando
 *											   los retornos de estas.
-------------------------------------------------------------------------- */

package main

import (
	"fmt"
	"net/http"
	"time"
)

// Revisar estado del servidor
func checkServer(url string, channel chan string) {
	_, error := http.Get(url)

	if error == nil {
		channel <- "¡Servidor " + url + " disponible!"
	} else {
		channel <- "¡Servidor " + url + " no disponible!"
	}
}

func main() {
	start := time.Now()

	channel := make(chan string) // Canal.

	servers := []string{
		"https://oregoom.com/",
		"https://udemy.com/",
		"https://google.com/",
		"https://youtube.com/",
	}

	for _, v := range servers {
		go checkServer(v, channel) // Concurrencia.
	}

	// Mostrar comunicación del canal
	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}

	processTime := time.Since(start)
	fmt.Println(processTime)
}
