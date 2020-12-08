package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS   \t", runtime.GOOS) //constant string, dont need parentheses  to call
	fmt.Println("GOARCH \t", runtime.GOARCH)
	fmt.Println("CPUs  \t", runtime.NumCPU())         //these are func(),take no parameter,return int
	fmt.Println("GoRoutines", runtime.NumGoroutine()) //NumCPU and NumGoroutine, dont add s
	go foo()
	bar()
	fmt.Println("CPUs  \t", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())

}
func foo() {
	for i := 0; i < 6; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i := 0; i < 6; i++ {
		fmt.Println("bar:", i)
	}
}
