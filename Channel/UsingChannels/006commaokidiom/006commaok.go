package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(eve, odd, quit)
	receive(eve, odd, quit)
	fmt.Println("About to exit")
}
func send(eve, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}
func receive(eve, odd, quit <-chan int) {
	for {
		select {
		case v := <-eve:
			fmt.Println("value from eve", v)
		case v := <-odd:
			fmt.Println("value from odd", v)
		case value, ok := <-quit:
			if !ok {
				fmt.Println("from quit value and ok", value, ok)
				return //when didnt return, code keeps running infinitely
			} else {
				fmt.Println("from quit  ok", ok)
			}
		}
	}
}
