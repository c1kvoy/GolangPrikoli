package shapes

import (
	"math"
)

type Circle struct { 
	CentralPoint Point
	Radius float64 
}

type Point struct {
	X float64 
	Y float64
	Z float64
}

type Triangle struct {
	P1 Point
	P2 Point
	P3 Point
}


func (self Point) Vector(other Point) Point {
	return Point{self.X - other.X, self.Y - other.Y, self.Z - other.Z}
}

func (self Point) Length() float64 {
	return math.Sqrt(self.X*self.X + self.Y*self.Y + self.Z*self.Z)
}

func (self Point) VecMul(other Point) Point {
	return Point{
		self.Y*other.Z - self.Z*other.Y,
		self.Z*other.X - self.X*other.Z,
		self.X*other.Y - self.Y*other.X,
	}
}

func (t Triangle) Area() float64 {
	v1 := t.P2.Vector(t.P1)
	v2 := t.P3.Vector(t.P1)
	if v1.Length() == 0 || v2.Length() == 0 {
		return 0
	}
	return v1.VecMul(v2).Length() / 2
}

func (t Triangle) Perimetr() float64 { 
	v1 := t.P2.Vector(t.P1).Length()
	v2 := t.P3.Vector(t.P1).Length()
	v3 := t.P2.Vector(t.P3).Length()
	return v1 + v2 + v3
}

  
func (self Circle) Area() float64{	
	return math.Pi * math.Pow(self.Radius, 2) 
}

func  (self Circle) Perimetr() float64{
	return math.Pi * self.Radius * 2
}