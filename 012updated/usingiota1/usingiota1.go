package main

import (
	"fmt"
)

func main() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Printf("%d\t\t %b\n", kb, kb)
	fmt.Printf("%d\t\t %b\n", mb, mb) //the binary shows that mb is kb shifted left ten times
	fmt.Printf("%d\t %b\n", gb, gb)   //gb is kb shifted left 20 times or mb shifted 10 times
}
