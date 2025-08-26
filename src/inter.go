package main

import "fmt"

type cuadrado struct {
	lado float64
}

func (c cuadrado) area() float64 {
	return c.lado * 2
}

type rectangule struct {
	lado  float64
	ancho float64
}

func (r rectangule) area() float64 {
	return r.lado * r.ancho
}

type shape interface {
	area() float64
}

func calcularArea(s shape) {
	fmt.Println("Area", s.area())
}

func inter() {
	myCuadrado := cuadrado{lado: 10}
	myRectangule := rectangule{lado: 10, ancho: 5}

	fmt.Println("Area cuadrado:", myCuadrado.area())
	fmt.Println("Area rectangulo:", myRectangule.area())

	calcularArea(myCuadrado)
	calcularArea(myRectangule)
}
