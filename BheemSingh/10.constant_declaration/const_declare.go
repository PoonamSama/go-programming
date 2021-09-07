package main

import (
	"fmt"
)

// constant cannot declare as short declaration operator
const a = 12
const b = 22

const (
	c = 25
	d = 26
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
