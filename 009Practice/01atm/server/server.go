package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	var balance int64 = 9563
	ln, err := net.Listen("tcp", ":80")
	ercheck(err)
	defer ln.Close()

	connx, err := ln.Accept()
	ercheck(err)
	defer connx.Close()
	for {
		netdata, err := bufio.NewReader(connx).ReadString('\n') //gives string and error
		ercheck(err)
		if strings.TrimSpace(string(netdata)) == "n" {
			fmt.Println("TCP Server exiting...")
			return
		}
		y := strings.TrimSpace(string(netdata))
		fmt.Println("->", netdata)
		x, amount := amt(y, balance)
		//connx.Write([]byte(convert the thing into slice of bytes))
		switch x {
		case 0:
			connx.Write([]byte("Error!Please enter an integer"))
			break

		case 10:
			connx.Write([]byte("Error!Please enter a Positive value"))
			return
		case 20:
			connx.Write([]byte("Error!Please enter amount divisible by 100 "))
			return
		case 30:
			connx.Write([]byte("Error! Please enter amount less than or equal to 5000"))
			return
		case 40:
			connx.Write([]byte("Error!Please enter amount less than or equal to your account balance that is:\n"))
			return
		default:
			//count++
			connx.Write([]byte("You want to withdraw:"))
			fmt.Println("Your account balance:", balance)
			//denoms(amount)
			balance = balance - amount
			connx.Write([]byte("TRANSACTION SUCCESSFUL.YOUR BALANCE NOW IS:"))
			//connx.Write([]byte(balance))
			connx.Write([]byte("------------------------"))
			return
		}
		/*if count >= 5 {
			e := fmt.Errorf("Error! you have made maximum number of transactions for today: 5. Initiating Exit")
			log.Fatalln(e)
		}*/
		if balance < 100 {
			fmt.Println("Your total transactions today:") //,count
			e := fmt.Errorf("Error!your account balance is %v .It is less than 100.You can't withdraw", balance)
			log.Fatalln((e))
		}
	}
}
func ercheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
	return
}

func amt(y string, z int64) (int64, int64) {
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner:=y
	//fmt.Println("Enter the amount you want to withdraw:")
	//scanner.Scan()

	//input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	input, err := strconv.ParseInt(y, 10, 64)
	fmt.Println(input)
	if err != nil {
		fmt.Println(err)
		return 0, input
	}
	if input <= 0 {
		return 10, input
	}
	if input%100 != 0 {
		return 20, input
	}
	if input > 5000 {
		return 30, input
	}
	if z-input < 0 {
		return 40, input
	}
	return 50, input
}
