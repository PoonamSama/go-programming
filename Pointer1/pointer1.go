package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println("value of a :", a)
	fmt.Println(&a)
	b := &a
	fmt.Println(b)
	fmt.Println(*&a)
	c := &b
	fmt.Printf("%T\n", c)
	fmt.Println(*&b)
	fmt.Println(&b)
	fmt.Println(*&c)
	*b = 95
	fmt.Println("value of a now:", a)
}
