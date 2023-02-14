package main

import "fmt"

func main() {
	a := 2

	a += 2
	fmt.Println(a)

	a -= 2
	fmt.Println(a)

	a *= 4
	fmt.Println(a)

	a /= 2
	fmt.Println(a)

	a %= 2
	fmt.Println(a)
}
