package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10
	y := 10
	c := generate(x, y)
	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}
}
func generate(x, y int) <-chan int {

	c := make(chan int)
	for i := 0; i < x; i++ {
		go func() { //means launch a gortn
			for j := 0; j < y; j++ { //this loop puts value in channel

				c <- j
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine()) //to see that we launched 1o grtn
	}
	return c
}
