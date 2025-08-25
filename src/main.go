package main

import (
	"fmt"
	"math"
)

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.16

	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	//Declaracion de varaibles
	base := 12          //Primera forma
	var altura int = 14 //Segunda forma
	var area int        //Go no compila si las variables no son usadas

	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna vaalores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)

	//Constant Declaration
	var height int = 14

	//Area Exercise
	base = 24
	height = 12
	fmt.Println("Rectangule area:", base*height)

	upperBase := 20
	fmt.Println("Trapeze Area: ", ((upperBase+base)*height)/2)

	var radius float64 = 5
	fmt.Println("Circle Area: ", pi2*math.Pow(radius, 2))

	fmt.Printf("Circulo %.2f \n", areaCirculo(2))
	fmt.Printf("Rectangulo %.2f \n", areaRectangulo(5, 10))
	fmt.Printf("Trapezoide %.2f \n", areaTrapezoide(10, 5, 3))

	var numeronose int
	var errorPrintln error
	numeronose, errorPrintln = fmt.Println("uno", "dos", "uno")
	fmt.Println(numeronose)
	fmt.Println(errorPrintln)

}

func areaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}
func areaRectangulo(base float64, altura float64) float64 {
	return base * altura
}

func areaTrapezoide(B float64, b float64, h float64) float64 {
	return h * (B + b) / 2
}
