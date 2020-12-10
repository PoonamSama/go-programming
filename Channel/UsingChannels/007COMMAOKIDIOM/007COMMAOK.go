package main

import "fmt"

func main() {
	c := make(chan int)
	d := make(chan int)
	go func() {
		c <- 48 //didnt close this channel
	}()
	value, ok := <-c
	fmt.Println(value, ok)
	//value, ok = <-c     //cant do more than once.error:deadlock all goroutines asleep
	//fmt.Println(value, ok)

	go func() {
		d <- 111
		close(d)
	}()
	dvalue, d_ok := <-d
	fmt.Println(dvalue, d_ok)
	dvalue, d_ok = <-d //DONT DO D_OK:=--- AGAIN BCZ THESE ARE NOT NEW VARIABLES
	fmt.Println(dvalue, d_ok)
	dvalue, d_ok = <-d //can do this any number of times
	fmt.Println(dvalue, d_ok)
	value, d_ok = <-d //can do this any number of times
	fmt.Println(dvalue, d_ok)

}
