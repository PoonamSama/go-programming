package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	ercheck(err)
	defer ln.Close()
	connx, err := ln.Accept()
	ercheck(err)
	for {
		netData, err := bufio.NewReader(connx).ReadString('\n')
		ercheck(err)
		fmt.Println(string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		connx.Write([]byte(myTime))

	}
	//connx.Close()
}
func ercheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
