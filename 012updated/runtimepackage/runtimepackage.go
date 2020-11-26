package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)   //go operating system
	fmt.Println(runtime.GOARCH) // amd64 means all int are int64 by default
}
