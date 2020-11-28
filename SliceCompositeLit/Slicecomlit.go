package main

import (
	"fmt"
)

func main() {
	// Composite Literal x:=type{values}
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	y := []string{"h", "ok", "ko", "hi", "ij"}
	fmt.Println(y)

}
