/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Es este archivo se demuestra la inicialización de variables que
 *		   solo pueden ser leídas dentro del scope de la condicional que las
 *		   inicializo.
-------------------------------------------------------------------------- */

package main

import "fmt"

func main() {
	if name, age := "Lucas", 26; name == "Lucas" {
		fmt.Println("Hola", name)
	} else {
		fmt.Printf("• Nombre: %s | • Edad: %d", name, age)
	}

	users := make(map[string]string)
	users["Lucas"] = "hozlucas28@gmail.com"
	users["Nahuel"] = "hoznahuel28@gmail.com"

	/*
		email, exist := users["Lucas"]
		fmt.Println(email, exist)
	*/

	if email, exist := users["Lucas"]; exist {
		fmt.Println("• Email:", email)
	} else {
		fmt.Printf("• Email: %s | • ¿Existe?: %t", email, exist)
	}
}
