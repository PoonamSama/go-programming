package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
}
func foo() {
	fmt.Println("defer is not done when os.exit() is called. this wont print!")
}

/*
... the Fatal functions call os.Exit(1) after writing the log message ...
*/

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
