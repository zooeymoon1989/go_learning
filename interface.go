package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1, x2, y1, y2 float64
}

type Shape interface {
	area() float64
	distance() float64
}

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) distance() float64 {

	a := r.x2 - r.x1
	b := r.y2 - r.y1

	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

}

func (s *Rectangle) area() float64 {
	l := s.x2 - s.x1
	w := s.y2 - s.y1

	return l * w
}

func main() {

	r := Rectangle{1, 2, 3, 5}

	fmt.Println(r.area())

	fmt.Println(r.distance())

	c := Circle{12}

	fmt.Println(c.area())
}
