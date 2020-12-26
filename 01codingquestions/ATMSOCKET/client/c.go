package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	c, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Your input>> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Println("->: " + message)
		if message == "Maximum transactions reached for a day:5 " {
			fmt.Println("Initiating exit")
			return
		}
		if message == "Balance less than 100.You can't withdraw " {
			fmt.Println("Initiating exit")
			return
		}
		if message == "exitnow" {
			fmt.Println("You chose no;exit")
			return
		}
	}
}
