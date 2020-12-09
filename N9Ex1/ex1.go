package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("goroutines:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("hi frm one")
		runtime.Gosched()
		fmt.Println("goroutines:", runtime.NumGoroutine())
		wg.Done()
	}()
	go func() {
		fmt.Println("hi frm two")
		fmt.Println("goroutines:", runtime.NumGoroutine())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("does this print?: unpredictability in concurrency")
	fmt.Println("goroutines:", runtime.NumGoroutine())
}
