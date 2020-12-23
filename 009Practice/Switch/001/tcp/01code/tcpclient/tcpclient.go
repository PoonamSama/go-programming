package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connx, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer connx.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(connx, text)

		message, _ := bufio.NewReader(connx).ReadString('\n')
		fmt.Print("->: " + message)
	}
}
