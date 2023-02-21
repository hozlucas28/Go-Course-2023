package main

import "figures/figures"

/* --------------------------- Programa Principal --------------------------- */

func main() {
	square := figures.Square{Side: 4}
	circle := figures.Circle{Radio: 5}

	figures.FigureLenghts(&square)
	figures.FigureLenghts(&circle)
}
