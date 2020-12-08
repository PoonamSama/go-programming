package main

import (
	"fmt"
)

func doSomething(x int) int { //this func is supposed to return a value

	return x * 5 //but when goroutine is called,return VALUES ARE DISCARDED
}

func main() {
	ch := make(chan int)
	go func() { //new goroutine(concurrent code)
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}
