package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("GoRoutine", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutine", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutine", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
