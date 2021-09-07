package main

import "fmt"

func main() {

	bs := []byte{71, 111}
	fmt.Printf("%s\n", bs)
	// above is showing string is aaray of bytes.

	s1 := "Hi, we are student of go."
	bs1 := []byte(s1)
	s2 := "there is no stop in go."
	bs2 := []byte(s2)
	fmt.Printf("%s\n%d\n%s\n%d\n", bs1, bs1, bs2, bs2)

}
