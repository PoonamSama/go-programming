package main

import "fmt"

func main() {

	c := make(<-chan int, 2) //this channel is type receive only. we can only send values to it
	go func() {
		c <- 42 //cant send value to rx only channel. invalid operation
		c <- 45
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T", c)
}
