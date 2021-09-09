package main

import "fmt"

func main() {
	fmt.Println("hello world, I am starting golang learning")
	controlflow()
	evennum()
}
func controlflow() {
	fmt.Println("control flow is how execution of statements are going. ")
	fmt.Println("simplest control flow is sequence flow. ")
}
func evennum() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
