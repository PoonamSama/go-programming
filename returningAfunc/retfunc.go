package main

import (
	"fmt"
)

func main() {
	s1, i1 := foo()
	fmt.Println(s1, i1)

}
func foo() (string, int) {
	s := "Poonam"
	x := 42
	return s, x
}
