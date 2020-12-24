package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

var db *sql.DB

func main() {

	connx, err := net.Dial("tcp", ":80")
	ercheck(err)
	defer connx.Close()

	for {
		input := bufio.NewReader(os.Stdin)
		fmt.Println(">>give your input")
		text, _ := input.ReadString('\n')
		fmt.Fprintf(connx, text+"\n")
		result, err := ioutil.ReadAll(connx)
		ercheck(err)
		fmt.Println(string(result))

		msg, _ := bufio.NewReader(connx).ReadString('\n')
		fmt.Println("->" + msg)
		if strings.TrimSpace(string(text)) == "n" {
			fmt.Println("TCP Client exiting...")
			return
		}
	}
}
func ercheck(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

}
