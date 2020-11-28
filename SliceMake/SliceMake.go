package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 100) // type, length and capacity
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) //capacity
	x[7] = 88           //we can do this
	//we cant do this:x[10]=5 it will give index out of length error
	x = append(x, 5) //this will assign 5 to x[10],also increase its length
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println("Slice y:")
	y := make([]int, 10, 12) // type, length and capacity
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
	y = append(y, 5)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
	y = append(y, 555)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	y = append(y, 99)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y)) //capacity changed , threw away the old array

}
