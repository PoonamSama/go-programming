package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS) //go operating system
	fmt.Println(runtime.GOARCH)
}
