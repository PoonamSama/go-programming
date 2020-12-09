package main

import "fmt"

func main() {

	c := make(chan int, 1) //we are making a buffer that lets ONLY 1 value to stay
	c <- 42
	c <- 45 //buffer only has 1 space, not two, hence we will get deadlock error
	fmt.Println(<-c)
}
