package main

import (
	"first/interfaces/shapes"
	"fmt"
)

type Shape interface {
	Area() float64
	Perimetr() float64
}

func AreaOut(figures []Shape) {
	for _, figure := range figures {
		fmt.Printf("%.2f \n", figure.Area())
	}
}

func main() {

	point1 := shapes.Point{X: 0.0, Y: 0.0, Z: 4.0}
	point2 := shapes.Point{X: 3.0, Y: 0.0, Z: 0.0}
	point3 := shapes.Point{X: 3.0, Y: 4.0, Z: 0.0}

	var triangle1 Shape = shapes.Triangle{
		P1: point1,
		P2: point2,
		P3: point3,
	}
	var circle1 Shape = shapes.Circle{
		Radius: 3,
	}
	figures := []Shape{circle1, triangle1}
	AreaOut(figures)
}
