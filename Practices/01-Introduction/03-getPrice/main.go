package main

import "fmt"

func getIGV(price float32) float32 {
	return price * 0.18
}

func getPriceOfSell(price, igv float32) float32 {
	return price + igv
}

func main() {
	var price float32

	fmt.Println("• Ingrese el valor de venta del producto:")
	fmt.Scanln(&price)

	var igv float32 = getIGV(price)
	var priceOfSell float32 = getPriceOfSell(price, igv)

	fmt.Printf("EL PRECIO DE VENTA ES $%0.2f Y EL IGV ES $%0.2f.", priceOfSell, igv)
}
