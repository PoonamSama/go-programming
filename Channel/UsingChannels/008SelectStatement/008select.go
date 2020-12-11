package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	d := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		c <- "Audio Message from c"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		d <- "Video Message from d"
	}()
	receive(c, d)
}
func receive(c, d <-chan string) {
	for i := 0; i < 2; i++ {
		select {
		case v := <-c:
			fmt.Println("Message:", v)
		case v := <-d:
			fmt.Println("Message:", v)
		}
	}
}
