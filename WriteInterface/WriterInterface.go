package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello from os")
	io.WriteString(os.Stdout, "Hello from io")
}
