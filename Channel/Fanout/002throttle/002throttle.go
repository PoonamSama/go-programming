package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go load(c1)
	go fanoutin(c1, c2)
	for v := range c2 {
		fmt.Println(v)
	}
}
func load(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
func fanoutin(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10 //we dnt want to range over c1 to know how many gortn we want to launch
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- somework(v2)
				}(v)
				wg.Done()

			}
		}()
	}
	wg.Wait()
	close(c2)
}
func somework(y int) int {
	return y + 100
}
