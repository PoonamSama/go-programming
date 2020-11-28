package main

import (
	"fmt"
)

func main() {
	x := []string{"Apple", "Mac", "IPad", "Iball"}
	y := []string{"Sam", "Sung", "Korea", "Future"}
	z := [][]string{x, y}

	fmt.Println(z)
}
