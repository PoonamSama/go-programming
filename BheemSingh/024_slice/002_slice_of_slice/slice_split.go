package main

import "fmt"

func main() {

	x := []int{3, 4, 5, 9, 11, 55, 87, 23, 54, 65, 75}

	a := x[:5]

	b := x[2:6]

	c := x[3:]

	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
