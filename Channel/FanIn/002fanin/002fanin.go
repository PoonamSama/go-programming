package main

import (
	"fmt"
)

func main() {
	x := fanIn(intro("James"), intro("Jessi"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
	fmt.Println("Thats it")
}
func intro(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("my name is %s and number: %d", name, i) //cant send a value in the same goroutine,hence go func
		}
	}()
	return c
}
func fanIn(label1, label2 <-chan string) <-chan string {
	m := make(chan string)
	go func() {
		for {
			m <- <-label1
		}
	}()
	go func() {
		for {
			m <- <-label2
		}
	}()
	return m
}
