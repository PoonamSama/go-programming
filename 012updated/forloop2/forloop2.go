package main

import (
	"fmt"
)

func main() {
	x, y := 1, 2
	a, b := x+5, y-1
	fmt.Println(x, y)
	fmt.Println(a, b)
	s := "Hello Aloha"
	bs := []byte(s)
	fmt.Println(bs)
	// reversing a string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	fmt.Println(bs)
	fmt.Printf("%c\n", bs)
}
