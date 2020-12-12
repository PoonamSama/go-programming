package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Create("no-files.txt") //returns a file ptr an error
	if err != nil {
		fmt.Println("theres an error", err)
	}
}
