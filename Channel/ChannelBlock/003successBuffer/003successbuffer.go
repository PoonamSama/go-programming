package main

import "fmt"

func main() {

	c := make(chan int, 1) //we are making a buffer that lets 1 value to stay
	c <- 42
	fmt.Println(<-c)
}
