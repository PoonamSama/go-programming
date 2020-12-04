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

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c square) area() float64 {
	return c.side * c.side
}
func main() {
	c1 := circle{
		radius: 2,
	}
	s1 := square{
		side: 4,
	}
	x1 := c1.area()
	y1 := s1.area()
	fmt.Println(x1, y1)
}
