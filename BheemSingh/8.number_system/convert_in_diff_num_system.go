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
	//feedback : it would be more informative to also print what you're intending to print for example
	fmt.Printf("Binary values: %b\n", bs)

}
