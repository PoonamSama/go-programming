package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var balance int64 = 9563

func main() {
	count := 0
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		y := strings.TrimSpace(string(netData))
		fmt.Println("we got from client-> ", string(netData))
		if count >= 5 {
			fmt.Println("Error! you have made maximum number of transactions for today: 5. Initiating Exit")
			c.Write([]byte("Maximum transactions reached for a day:5 "))
			return
		}
		if balance < 100 {
			c.Write([]byte("Balance less than 100.You can't withdraw "))
			return
		}
		x, amount := amt(y, balance)
		switch x {
		case 0:
			fmt.Println("Error!Please enter an integer")
			c.Write([]byte("Error!Please enter an integer \t"))
			break

		case 10:
			fmt.Println("Error!Please enter a Positive value")
			c.Write([]byte("Error!Please enter a Positive value \t"))
			break
		case 20:
			fmt.Println("Error!Please enter amount divisible by 100 ")
			c.Write([]byte("Error!Please enter amount divisible by 100 \t"))
			break
		case 30:
			fmt.Println("Error! Please enter amount less than or equal to 5000")
			c.Write([]byte("Error! Please enter amount less than or equal to 5000 \t"))
			break
		case 40:
			fmt.Printf("Error!Please enter amount less than or equal to your account balance that is: %v\n", balance)
			j := fmt.Sprintf("Error!Please enter amount less than or equal to your account balance that is: %v\t", balance)
			c.Write([]byte(j))
			break
		default:
			count++
			i := denoms(amount)
			fmt.Println(i)
			balance = balance - amount
			fmt.Println("TRANSACTION SUCCESSFUL.YOUR BALANCE NOW IS:", balance)
			j := fmt.Sprintf(" \t TRANSACTION SUCCESSFUL.YOUR BALANCE NOW IS: %d\t ", balance)
			c.Write([]byte(i))
			c.Write([]byte(j))
			break
		}
		c.Write([]byte("Do you want to start another transaction?Enter y to continue and n to exit"))
		newline := "\n"
		c.Write([]byte(newline))
		foo(c)
		newline = "\n"
		c.Write([]byte(newline))
	}

}
func amt(y string, balance int64) (int64, int64) {

	input, err := strconv.ParseInt(y, 10, 64)
	if err != nil {
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
	if balance-input < 0 {
		return 40, input
	}
	return 50, input
}

func foo(c net.Conn) {

	fData, ferr := bufio.NewReader(c).ReadString('\n')
	if ferr != nil {
		fmt.Println(ferr)
		return
	}
	ln := strings.TrimSpace(string(fData))

	switch ln {
	default:
		{
			c.Write([]byte("Server says:Error!Wrong input.You can only Enter y or n"))
			newline := "\n"
			c.Write([]byte(newline))
			foo(c)
			return
		}

	case "y":
		{
			c.Write([]byte("You chose:yes, start a new transaction.Please enter the amount"))
			return
		}
	case "n":
		{
			c.Write([]byte("exitnow"))
			c.Close()
		}
	}
}

func denoms(x int64) string {
	var q1 int64
	var q2 int64
	var q3 int64
	q1 = x / 500
	x = x % 500
	q2 = x / 200
	x = x % 200
	q3 = x / 100
	fmt.Printf("500*%d +200*%d +100*%d \n", q1, q2, q3)
	i := fmt.Sprintf("500*%d +200*%d +100*%d ", q1, q2, q3)
	return i
}
