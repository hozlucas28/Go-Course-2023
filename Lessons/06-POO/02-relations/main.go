/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   En las relaciones no se pueden acceder a los atributos de la
 *		   herencia directamente, pero si a través de la relación creada.
-------------------------------------------------------------------------- */

package main

import "fmt"

/* --------------------------- Relación Uno A Uno --------------------------- */

type User struct {
	name   string
	email  string
	active bool
}

type Student struct {
	user User // Relación estudiante --> usuario.
	code string
}

/* -------------------------- Relación Uno A Muchos ------------------------- */

type Video struct {
	course Course // Relación videos --> curso.
	title  string
}

type Course struct {
	title  string
	videos []Video
}

/* --------------------------- Programa Principal --------------------------- */

func main() {
	// Relación de uno a uno
	/*
		user1 := User{
			name:   "Lucas",
			email:  "lucas@gmail.com",
			active: true,
		}

		user2 := User{
			name:   "Juan",
			email:  "juan@gmail.com",
			active: true,
		}

		student1 := Student{
			user: user1,
			code: "6X616DF65HDJ",
		}

		fmt.Println(user2)    // Usuario no estudiante
		fmt.Println(student1) // Usuario estudiante
	*/

	// Relación de uno a muchos
	video1 := Video{title: "01 Video"}
	video2 := Video{title: "02 Video"}
	video3 := Video{title: "03 Video"}

	course := Course{
		title:  "Curso de Go",
		videos: []Video{video1, video2, video3},
	}

	video1.course = course
	video2.course = course
	video3.course = course

	for _, v := range course.videos {
		fmt.Println(v)
	}
}
