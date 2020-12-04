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

func (c circle) area() float64 {     //method area
	return math.Pi * c.radius * c.radius
}
func (c square) area() float64 {   //method area
	return c.side * c.side
}
type shape interface{
	area() float64
}
func info (s shape){
	fmt.Println(s.area())
}
func main() {
	c1 := circle{
		radius: 2,
	}
	s1 := square{
		side: 4,
	}
	info (c1)
	info (s1) 
}
