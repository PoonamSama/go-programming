package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go loadchannel(c1)
	go fanoutin(c1, c2)
	for v := range c2 {
		fmt.Println(v)
	}
}
func loadchannel(c chan int) {
	for i := 0; i < 20; i++ {
		c <- i
	}
	close(c)
}
func fanoutin(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)        //Add just before launching goroutine
		go func(x int) { //x int is parameter for the anon func
			c2 <- somework(x) //launching 20 goroutines to do somework() this is fanout and sending these values to c2 channel is fanin
			wg.Done()
		}(v) //we passed v from range c1 as arg to the anon func

	}
	//close(c2) //this gave err panic:send to closed channel
	wg.Wait()
	close(c2)
}
func somework(y int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return y + 1000 //this value is returned to c2<-somework(returned val)
}
