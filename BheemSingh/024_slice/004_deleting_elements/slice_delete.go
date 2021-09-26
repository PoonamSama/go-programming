package main

import (
	"fmt"
)

func main() {

	x := []int{42, 55, 66, 76, 78}
	y := []int{22, 12, 34, 54, 65, 12}

	x = append(x, y...)
	fmt.Println(x)

	// if we don't want 22 and 12

	x = append(x[:5], x[7:]...)
	fmt.Println(x)

}
