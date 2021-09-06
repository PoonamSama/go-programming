package main

import "fmt"

var a int

type GOLANG int

var b GOLANG

func main() {

	a = 42
	b = 1

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
