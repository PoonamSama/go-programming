package main

import "fmt"

func main() {

	x := sum(2, 3, 4, 5, 6, 7, 8)

	fmt.Printf("\n Sum of all numbers is %d \n ", x)
}

func sum(x ...int) int {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("for index value %d , we are adding %d and sum is %d \n", i, v, sum)

	}
	return sum
}
