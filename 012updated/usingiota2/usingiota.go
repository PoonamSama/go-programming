package main

import (
	"fmt"
)

const (
	_  = iota               // _to eliminate ;iota=0 itself
	kb = (1 << (iota * 10)) //kb=1024
	mb = (1 << (iota * 10)) //mb is 1 <<20, iota will increment to 2 itself
	gb = (1 << (iota * 10)) //gb is 1<<30
)

func main() {

	fmt.Printf("%d\t\t %b\n", kb, kb)
	fmt.Printf("%d\t\t %b\n", mb, mb) //the binary shows that mb is kb shifted left ten times
	fmt.Printf("%d\t %b\n", gb, gb)   //gb is kb shifted left 20 times or mb shifted 10 times
}
