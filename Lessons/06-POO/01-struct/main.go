/* --------------------------------------------------------------------------
 * IMPORTANTE:
 *  			  - La asignación de valores (segundo método de inicialización)
 *					de los atributos de un objeto solo se puede utilizar en los
 *					propios y no en los heredados.
-------------------------------------------------------------------------- */

package main

import "fmt"

// Struct (objeto) persona
type Person struct {
	name string
	age  int
}

// Métodos
func (p *Person) showAttributes() {
	fmt.Printf("nombre: %s\nedad: %d\n\n", p.name, p.age)
}

func (p *Person) modifyName(newName string) {
	p.name = newName
}

// Herencia
type Employeer struct {
	Person
	salary float32
}

func main() {
	// Inicialización de un objeto - Primer método
	person1 := Person{"Lucas", 20}
	person1.showAttributes()

	// Modificación de un atributo
	// person1.name = "Nahuel"
	person1.modifyName("Nahuel")
	person1.showAttributes()

	// Inicialización de un objeto - Segundo método
	person2 := Person{name: "Juan", age: 32}
	person2.showAttributes()

	employeer1 := Employeer{salary: 500}
	employeer1.name = "Pedro"
	employeer1.age = 46
	employeer1.showAttributes()
	fmt.Println(employeer1)
}
