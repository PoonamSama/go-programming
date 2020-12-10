package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 50; i++ {
			c <- i
		}
		close(c) //if you dnt close the channel,it will keep waiting for more values BECZ OF (RANGE C) and
	}() //fatat error:all goroutines asleep,deadlock

	for v := range c {
		fmt.Println(v) //to range over a channel; CLOSING CHANNEL IS NECESSARY
	}
	fmt.Println("about to exit code")
}
