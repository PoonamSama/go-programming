package main

import "fmt"

func main() {
	x := []int{1, 2, 44, 32, 67}
	fmt.Println(x)
	fmt.Println(x[1])   // to access this element
	fmt.Println(x[1:])  //: does the slicing before index 1
	fmt.Println(x[2:3]) // does slicing on original x
}
