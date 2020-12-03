package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground,this is main")
	defer foo()
	bar()
	fmt.Println("main again")
	
}
func foo(){
fmt.Println("foo here")
}
func bar(){
fmt.Println("bar here")
}