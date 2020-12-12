package main

import (
	"fmt"
	"io"
	"strings"

	//	"log"
	"os"
)

func main() {
	f, err := os.Create("no-file.txt")
	if err != nil {
		fmt.Println("error occurred:", err)
		return
	}
	defer f.Close()
	r := strings.NewReader("some valueeeee")
	io.Copy(f, r)

}
