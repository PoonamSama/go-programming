package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var grogu string
var balance1 = 9563
var count = 0
var getbody string

func main() {
	fmt.Print("start main")
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	// read request
	request(conn)
	// write response
	respond(conn)
}

func request(conn net.Conn) {
	var x string //
	i, a := 0, 0
	loop := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			path := strings.Fields(ln)[1]
			fmt.Println("***METHOD", m)
			if m == "GET" {
				grogu = fmt.Sprint("balance")
			} else if m == "POST" && path == "/logout" {
				fmt.Println("Initiating exit")
				getbody = "Initiating exit"
				os.Exit(1)
			} else if m == "POST" {
				fmt.Println("starting transaction...")
				grogu = fmt.Sprint("startt")
			} else { //if any other method is given
				grogu = fmt.Sprint("InvalidMethod")
				getbody = "Invalid Method. Only GET and POST are acceptable."
			}
		}

		if ln == "" {
			// headers are done
			x = fmt.Sprint(i)
			loop++
			if loop > 1 {
				printing(grogu, balance1)
				break
			}
			continue
		}
		i++
		a, _ = strconv.Atoi(x)
		if a != 0 && i == a+1 {
			p := strings.Fields(ln)
			b := foo(p, grogu)
			if b == "stop" {
				break
			}

			break
		}

	}
}

func foo(p []string, x string) string {

	if x == "balance" {
		fmt.Println("Your balance is:", balance1)
		j := fmt.Sprintf("Your balance: %d\n", balance1)
		getbody = j
	} else if x == "startt" {
		if len(p) > 1 {
			if count >= 5 {
				getbody = "You have made maximum transactions for today:5. You can logout using POST/logout"
			} else {
				fmt.Println("You can only enter ONE INTEGER VALUE")
				getbody = "You can only enter ONE INTEGER VALUE"
			}
			return "stop"

		}
		if count >= 5 {
			fmt.Println("Error! you have made maximum number of transactions for today: 5.")
			getbody = "Maximum transactions reached for a day:5.  You can logout using POST/logout"
			return "stop"
		}
		if balance1 < 100 {
			getbody = "Balance less than 100.You can't withdraw.  You can logout using POST/logout"
			return "stop"
		}
		c, amount := transaction(p[0], balance1)
		switch c {
		case 0:
			fmt.Println("Error!Please enter an integer")
			getbody = "Error!Please enter an integer"
			break

		case 10:
			fmt.Println("Error!Please enter a Positive value")
			getbody = "Error!Please enter a Positive value"
			break
		case 20:
			fmt.Println("Error!Please enter amount divisible by 100 ")
			getbody = "Error!Please enter amount divisible by 100 "
			break
		case 30:
			fmt.Println("Error! Please enter amount less than or equal to 5000")
			getbody = "Error! Please enter amount less than or equal to 5000"
			break
		case 40:
			fmt.Printf("Error!Please enter amount less than or equal to your account balance that is: %v\n", balance1)
			j := fmt.Sprintf("Error!Please enter amount less than or equal to your account balance that is: %v\t", balance1)
			getbody = j
			break
		default:
			count++
			i := denoms(amount)
			balance1 = balance1 - amount
			fmt.Println("TRANSACTION SUCCESSFUL.YOUR BALANCE NOW IS:", balance1)
			j := fmt.Sprintf(" \n TRANSACTION SUCCESSFUL. Enter value and Send if you want to start another transaction.  You can logout using POST/logout")
			getbody = i + j
			break
		}
	} else {
		getbody = "Invalid method. Only GET and POST are acceptable."
		return "stop"
	}
	return "ok"
}
func respond(conn net.Conn) {
	body1 := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello there World</strong></body></html>`
	//body := getbody
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body1))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body1)
}
func printing(grogu string, balance1 int) {
	fmt.Println("count is:", count)
	fmt.Println("client didnt give a value or body")
	if grogu == "balance" {
		fmt.Println("method get no input, the balance is:", balance1)
		j := fmt.Sprintf("Your balance: %d\n", balance1)
		getbody = j
	}
	if grogu == "startt" {
		if count >= 5 {
			fmt.Println("Error! you have made maximum number of transactions for today:5")
			getbody = "Maximum transactions reached for a day:5"

		} else {
			fmt.Println("client must give some value")
			getbody = "Please enter a value. Field can't be empty."
		}
	}
}
func transaction(p string, bal int) (int, int) {
	input, err := strconv.Atoi(p)
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
	if bal-input < 0 {
		return 40, input
	}

	return 50, input
}
func denoms(x int) string {
	var q1 int
	var q2 int
	var q3 int
	q1 = x / 500
	x = x % 500
	q2 = x / 200
	x = x % 200
	q3 = x / 100
	fmt.Printf("500*%d +200*%d +100*%d \n", q1, q2, q3)
	i := fmt.Sprintf("500*%d +200*%d +100*%d ", q1, q2, q3)
	return i
}
