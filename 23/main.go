package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) GetDistance(to Point) float64 {
	f := p.x - to.x
	s := p.y - to.y

	result := f*f + s*s
	return math.Sqrt(result)
}

func main() {
	p := NewPoint(1.0, 1.0)
	s := NewPoint(3.0, 3.0)

	fmt.Println(p.GetDistance(s))
}