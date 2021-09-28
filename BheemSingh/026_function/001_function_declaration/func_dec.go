package main

import "fmt"

func main() {
	foo("Neha")
	fmt.Printf("Sum is : %d ", sum(20, 12))
}

func sum(a, b int) int {
	sum := a + b
	return sum
}

func foo(s string) {
	fmt.Printf("Hello, %s\n ", s)
}
