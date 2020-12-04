package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	side float64
}

func (s *circle) area() float64 {
	return math.Pi * s.radius * s.radius
}
func (s square) area() float64 {
	return s.side * s.side
}

type shape interface {
	area() float64
}

func bar(h1 shape) {
	fmt.Println(h1.area())
}
func main() {
	c1 := circle{
		radius: 2,
	}
	s1 := square{
		side: 2,
	}
	bar(&c1)
	bar(s1)

}
