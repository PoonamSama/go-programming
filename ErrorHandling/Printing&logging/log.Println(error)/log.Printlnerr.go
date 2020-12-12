package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("newfile.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	f2, err := os.Open("no-file.txt")

	if err != nil {
		log.Println("An Error has occurred", err) //This writes the msg in newfile.txt
	}
	defer f2.Close()
	fmt.Println("check the newfile.txt")

}
