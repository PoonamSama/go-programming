package main

import (
	"fmt"
)

func main() {
	s := ("Hello Playground")

	fmt.Println(s)
	bs := []byte(s) //slice of bytes or uint8
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i]) //gives the UTF8 code
	}
	fmt.Println("\n")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\t", s[i])
	}

}
