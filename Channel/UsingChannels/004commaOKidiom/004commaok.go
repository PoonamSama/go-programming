package main

import "fmt"

func main() {
	eve := make(chan int)
	od := make(chan int)
	qut := make(chan bool)
	go send(eve, od, qut)
	receive(eve, od, qut)
}
func send(eve, od chan<- int, qut chan<- bool) {
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			od <- i
		}

	}
	close(qut)
}
func receive(eve, od <-chan int, qut <-chan bool) {
	for {
		select {
		case v := <-eve:
			fmt.Println("from eve", v)
		case v := <-od:
			fmt.Println("from od", v)
		case i, ok := <-qut: //1st is value and 2nd is bool,if chan closed then value is 0 &bool is false
			if !ok { //bool ok is false since we close(qut)
				fmt.Println("from comma ok", i, ok)
				return
			} else {
				fmt.Println("from comma ok", i)
			}

		}
	}

}
