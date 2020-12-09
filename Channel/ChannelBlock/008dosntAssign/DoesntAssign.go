package main

import "fmt"

func main() {
	c := make(chan int)
	c1 := make(chan<- int) //send only channel. c1<-4 will work
	c2 := make(<-chan int) //receive only channel. fmt.Println(<-c2) will work
	fmt.Printf("type of c\t:%T\n", c)
	fmt.Printf("type of c1\t%T\n", c1)
	fmt.Printf("type of c2\t%T\n", c2)

	//	c2 = c //doesnt give err, but also doesnt change type
	//fmt.Printf("type of c2\t%T\n", c2)
	//c = c2  //err: cannot use c2 (type <-chan int) as type chan int in assignment
	//c2 = c1 // also gives error: cannot use c1 (type chan<- int) as type <-chan int in assignment
	//general to specific can  convert
	fmt.Printf("c: %T\n", (chan<- int)(c))
	//specific to general cant convert
	//fmt.Printf("c2: %T\n", (chan int)(c2)) //gives error
	//sp to sp cant convert
	fmt.Printf("c1: %T\n", (<-chan int)(c1)) //err cant convert

}
