package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("Anon func")
	}()
	func(x int, y string) {
		fmt.Println("anon func was passed:", x, y, "as arguments")
	}(25, "Lady")
}
func foo() {
	fmt.Println("foo ran")
}
