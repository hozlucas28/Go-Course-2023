package figures

import "fmt"

type iFigure interface {
	getArea() float32
	getPerimeter() float32
}

func FigureLenghts(iF iFigure) {
	fmt.Println("• Medidas:", iF)
	fmt.Println("• Area:", iF.getArea())
	fmt.Println("• Perímetro:", iF.getPerimeter())
	fmt.Printf("\n")
}
