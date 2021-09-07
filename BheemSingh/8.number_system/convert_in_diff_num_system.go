package main

import (
	"fmt"
)

func main() {

	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Printf("%b\n", bs)
	fmt.Printf("%#X ", bs)

}
