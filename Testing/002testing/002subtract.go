package main

import "fmt"

func main() {
	fmt.Println("7-5 is", sub(7, 5))
}
func sub(x, y int) int {
	c := x - y
	return c
}
