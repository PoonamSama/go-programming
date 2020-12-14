package main

import (
	"fmt"

	acdc "github.com/GoesToEleven/go-programming/Testing/004exampletest/01acdc"
)

func main() {
	z := acdc.Sum(6, 1, 1, 2)
	fmt.Println(z)
	fmt.Printf("type of z is : %T", z)
}
