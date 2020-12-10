package main

import "fmt"

func main() {

	evenchan := make(chan int)
	oddchan := make(chan int)
	quitchan := make(chan int)
	go sending(evenchan, oddchan, quitchan)
	receive(evenchan, oddchan, quitchan)
}
func receive(evenchan, oddchan, quitchan <-chan int) {
	for { //an infinite loop until it stops
		select {
		case v := <-evenchan:
			fmt.Println("from evenchan", v)
		case v := <-oddchan:
			fmt.Println("from oddchan", v)
		case v := <-quitchan:
			fmt.Println("from quitchan", v)

			return
		}
	}
}

func sending(evenchan, oddchan, quitchan chan<- int) { //we'are going to SEND TO THESE CHANNELS
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			evenchan <- i
		} else {
			oddchan <- i
		}
	}
	//close(evenchan)  if we close(),extra 0 are printed
	//close(oddchan)
	quitchan <- 0

}
