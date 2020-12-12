package main

import "fmt"

func main() {
	c := make(chan int)
	for j := 0; j < 10; j++ { // rpt this 10 times,i.e. (launch gortn and put values in it) 10 times
		//to launch a go rtn:use go func,then load values or whatever
		go func() { //launching a  new goroutine
			for i := 0; i < 10; i++ {
				c <- i //sending 0 to 9 values to channel c
			} //we are done putting nums on the channel

		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}
