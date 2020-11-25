package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {
	a = 19
	b = 20
	fmt.Printf("value of a is %v and type is %T\n", a, a)
	fmt.Printf("value of b is %v and type is %T\n", b, b)
	a = int(b)
	fmt.Printf("value of b is %v and type is %T\n", b, b)
	fmt.Printf("value of a is %v and type is %T\n", a, a)
}
