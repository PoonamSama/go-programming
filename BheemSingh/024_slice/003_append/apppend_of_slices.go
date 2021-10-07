package main

import (
	"fmt"
)

func main() {

	x := []int{4, 7, 6, 8, 9, 11, 55}

	x = append(x, 12, 14)
	fmt.Println(x)

	y := []int{16, 18, 20, 22}

	x = append(x, y...)
	fmt.Println(x)

}
