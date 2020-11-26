package main

import "fmt"

const (
	a        = 42 //untyped, known as constant of a kind, gives compiler a lil flexibility
	b        = 42.234
	s string = "hello world"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", s)
}
