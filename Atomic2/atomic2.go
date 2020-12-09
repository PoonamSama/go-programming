package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("GoRoutine", runtime.NumGoroutine())
	fmt.Println("--------")

	var counter int64
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter:", atomic.LoadInt64(&counter))
			runtime.Gosched()

			wg.Done()
		}()
		fmt.Println("GoRoutine", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("--------")
	fmt.Println("GoRoutine", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
