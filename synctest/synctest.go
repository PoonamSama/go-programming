package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t    ", runtime.GOOS)    //constant string
	fmt.Println("GO ARCH\t", runtime.GOARCH) // these are just constants, don't need parentheses()
	fmt.Println("CPUs\t   ", runtime.NumCPU()) //func that return int
	fmt.Println("GO ROUTINES\t", runtime.NumGoroutine()) //same

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("CPUs\t   ", runtime.NumCPU())
	fmt.Println("GO ROUTINES\t", runtime.NumGoroutine())

	wg.Wait()
}
func foo() {
	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
func bar() {
	fmt.Println("ki")
}
