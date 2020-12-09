package main

import "fmt"

func main() {

	c := make(chan<- int, 2) //this channel is type send only. we can only send values to it
	go func() {
		c <- 42
		c <- 45
	}()
	fmt.Println(<-c)
	fmt.Println(<-c) //cant receive value frm chan. invalid operation
	fmt.Printf("%T", c)
}
