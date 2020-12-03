package main

import (
	"fmt"
)

func main() {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("totalsum is:", sum(a1...))
	fmt.Println("even sum is:", evensum(sum, a1...))
}
func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func evensum(f func(x ...int) int, y ...int) int {
	var xi []int //xi is a slice
	for _, v1 := range y {
		if v1%2 == 0 {
			xi = append(xi, v1)
		}
	}
	total := f(xi...) //unfurling xi bcz argument of func sum is x ...int
	return total

}
