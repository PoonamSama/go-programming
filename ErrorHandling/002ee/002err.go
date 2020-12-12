package main

import (
	"fmt"
	"log"
	"os"
	//"io/iooutil"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("err")
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Create("no-file.txt")
	if err != nil {
		fmt.Println("err")
	}
	defer f2.Close()
	fmt.Println("check the log.txt file in the directory")

}
