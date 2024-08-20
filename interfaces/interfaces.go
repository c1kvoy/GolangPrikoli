package main

import (
	"first/interfaces/shapes"
	"fmt"
)

type Shape interface {
	Area() float64
	Perimetr() float64
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
	fmt.Printf("Triangle1 Area is: %.2f \n", triangle1.Area())
	fmt.Printf("Triangle1 Perimetr is: %.2f \n", triangle1.Perimetr())
	fmt.Printf("Circle1 Area is: %.2f \n", circle1.Area())
	fmt.Printf("Circle1 Perimetr is: %.2f \n", circle1.Perimetr())
}
