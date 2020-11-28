package main

import (
	"fmt"
)

func main() {
	x := []int{1, 34, 5, 65, 90}
	for index, value := range x {
		fmt.Println(index, value)
		//fmt.Printf("At position %d, the value is %d\n", index, value)
		//if above statement isn't commented,firstprintlnd doesn't work
	}
	fmt.Println("The End")
}
