package main

/*
	import (
		"packages/messages" // Paquete.
	)
*/

/*
	import (
		"fmt"
		"packages/models"
	)
*/

/*
	import "github.com/donvito/hellomod"
*/

import (
	figures "github.com/hozlucas28/Go-Module-Practice-2023"
)

func main() {
	// Modularización
	/*
		messages.Greet()
		messages.Print()
	*/

	// Encapsulamiento
	/*
		p1 := models.Person{}
		p1.Constructor("Lucas", 20)
		p1.ShowAttributes()

		p1.SetName("Nahuel")
		fmt.Println(p1.GetName())

		p1.SetAge(51)
		fmt.Println(p1.GetAge())
	*/

	// Importación de módulos de terceros
	/*
		hellomod.SayHello()
	*/

	// Importación de módulos de terceros - Propio
	square := figures.Square{Side: 4}
	circle := figures.Circle{Radio: 5}

	figures.FigureLenghts(&square)
	figures.FigureLenghts(&circle)
}
