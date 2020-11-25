package main

import "fmt"

//func main is the entry point
func main() {
	fmt.Println("hey everyone i am so grateful")
	foo()
	fmt.Println("something")
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

//control flow:sequence;loop or iterative;conditional : how a computer reads code
func foo() {
	fmt.Println("AYO HITMAN BANG")
}
func bar() {
	fmt.Println("and then we exited")
}
