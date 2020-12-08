package main

import (
	"fmt"
)

type square struct {
	side int
}
type rect struct {
	lenght  int
	breadth int
}

func (x square) area() int { //method area returns an int
	return x.side * x.side
}
func (x rect) area() int {
	return x.lenght * x.breadth

}

type shape interface {
	area() int //interface has a method area() attached that returns an int
}

func printing(s shape) int {
	return s.area()
}
func main() {
	s1 := square{2}
	s2 := square{3}
	r1 := rect{1, 3}
	r2 := rect{4, 5}
	fmt.Println(s1.area()) //BASIC WAY FOR A VALUE TO ACCESS A METHOD
	fmt.Println(s2.area())
	fmt.Println(r1.area())
	fmt.Println(r2.area())
	xy := printing(r1)
	fmt.Println("area of r1 throughfunc printing is:", xy)
	fmt.Println("area of s2 throughfunc printing is:", printing(s2))
}
