/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   El encapsulamiento trata de crear un constructor para modificar
 *		   atributos privados del objeto.
 *
 *
 * IMPORTANTE:
 *  			  - Los métodos <get> me dejan obtener atributos privados.
 *  			  - Los métodos <get> me permiten modificar atributos privados.
-------------------------------------------------------------------------- */

package models

import "fmt"

type Person struct {
	// Atributos privados
	name string
	age  int
}

// Constructor
func (p *Person) Constructor(name string, age int) {
	p.name = name
	p.age = age
}

/* --------------------------------- Setters -------------------------------- */

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetAge(age int) {
	p.age = age
}

/* --------------------------------- Getters -------------------------------- */

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetAge() int {
	return p.age
}

/* --------------------------------- Métodos -------------------------------- */

func (p *Person) ShowAttributes() {
	fmt.Println(p)
}
