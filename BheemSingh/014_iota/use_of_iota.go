package main

import (
	"fmt"
)

// iota represents succesive untyped integer

const (
	x = iota
	y = iota
	z = iota
)
const (
	_ = iota
	a
	b
	c
)

const (
	d = iota
	e
	f
	_
	h
)

func main() {
	fmt.Println(x, y, z)
	fmt.Println(a, b, c, d, e, f, h)

}
