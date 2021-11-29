package main

import (
	"fmt"
	"math"
)

func main() {
	shapes := []Shape{}
	shapes = append(shapes, &Square{10})
	shapes = append(shapes, &Circle{5.5})

	fmt.Println(SumArea(shapes))

}

type Square struct {
	Length float64
}

func (sq *Square) Area() float64 {
	return sq.Length * sq.Length
}

type Circle struct {
	Radius float64
}

func (cr *Circle) Area() float64 {
	return cr.Radius * cr.Radius * math.Pi
}

type Shape interface {
	Area() float64
}

func SumArea(shapes []Shape) float64 {
	result := 0.0

	for _, shape := range shapes {
		result += shape.Area()
	}

	return result
}
