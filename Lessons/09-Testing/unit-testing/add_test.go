/* --------------------------------------------------------------------------
 *	- <go test -coverprofile=coverage.out> = Crea un archivo con los porcentajes
 *										     de testeo de las funciones.
 *	- <go tool cover -func=coverage.out> = Lee y muestra los porcentajes, con detalle,
 *										   de las funciones testeadas.
 *	- <go tool cover -html=coverage.out> = Crea un HTML de los porcentajes, mostrando
 *										   que lineás han sido testeadas en dichas funciones.
 *
 *	- <go test -cpuprofile=cou.out> = Crea un archivo que muestra los tiempos que han
 *									  tardado las pruebas en ejecutarse.
 *	- <go tool pprof cou.out> = Lee y muestra el archivo de los tiempos.
 *								- <top> = Ordena de mayor a menor tiempo.
 *								- <list <FUNCTION>> = Otorga mayor información de los
 *													  tiempos de la función (argumento).
 *								- <pdf> = Crea un PDF del reporte.
-------------------------------------------------------------------------- */

package unittesting

import "testing"

/*
	func TestAdd(t *testing.T) {
		operation := Add(5, 5)

		if operation != 10 {
			t.Errorf("Se esperaba 10 se obtuvo %d", operation)
		}
	}
*/

/* ------------------ Pruebas Convencionales Mediante Tabla ----------------- */

func TestAdd(t *testing.T) {
	table := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, v := range table {
		total := Add(v.a, v.b)

		if total != v.c {
			t.Errorf("Se esperaba %d, se obtuvo %d", v.c, total)
		}
	}
}

func TestGetMax(t *testing.T) {
	table := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{5, 3, 5},
		{25, 41, 41},
	}

	for _, v := range table {
		max := GetMax(v.a, v.b)

		if max != v.c {
			t.Errorf("Se esperaba %d, se obtuvo %d", v.c, max)
		}
	}
}

func TestFibonacci(t *testing.T) {
	table := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, v := range table {
		fibo := Fibonacci(v.n)

		if fibo != v.r {
			t.Errorf("Se esperaba %d, se obtuvo %d", v.r, fibo)
		}
	}
}
