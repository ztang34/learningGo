package main

import (
	"fmt"
)

func main() {
	sq1, err := NewSquare(-3, -2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sq1)
	}

	sq2, _ := NewSquare(-2, -2, 2)
	fmt.Println(sq2.Area())
	sq2.Move(3, 4)
	fmt.Println(sq2)

}

type Point struct {
	X int
	Y int
}

func (p *Point) MovePoint(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length int
}

func (sq *Square) Move(dx int, dy int) {
	sq.Center.MovePoint(dx, dy)
}

func (sq *Square) Area() int {
	return sq.Length * sq.Length
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be a postive value")
	}

	return &Square{Point{x, y}, length}, nil
}
