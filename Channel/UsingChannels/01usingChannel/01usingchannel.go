package main

import "fmt"

func main() {
	c := make(chan int) // c is bidirectional channel
	//send
	go foo(c) //another goroutine starts
	//receive
	bar(c) //must remove go from here because we want to print it,and we're not waiting.
	//we dont want THIS BAR() to go to another routine bcz then it wont print w/o waitgrps
}

//send
func foo(c chan<- int) { //send only channel;we can convert bidi chan to unidi channel(gen to specific)
	c <- 42 //we SENT 42 to c;no need to WAIT for this,it happened in another goroutine
}

//receive
func bar(c <-chan int) { //rx only channel
	fmt.Println(<-c)

}
