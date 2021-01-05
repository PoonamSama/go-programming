package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var balance = 9563
var transactionCount = 0

//INPUT exported
type INPUT struct {
	AMOUNT int `json:"amount"`
}

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Your balance is": balance})
	})
	server.POST("/", func(c *gin.Context) {
		var input INPUT
		c.BindJSON(&input)
		c.JSON(200, gin.H{"Result": transaction(input.AMOUNT)})
	})
	server.POST("/exit", func(c *gin.Context) {
		c.JSON(200, "Exit")
		os.Exit(1)
	})
	server.Run(":8080") // listen for incoming connections
}
func transaction(clientInput int) string { //response func for method POST/
	if transactionCount >= 5 {
		fmt.Println("Error! you have made maximum number of transactions for today: 5.")
		message := fmt.Sprintf("Maximum transactions reached for a day:5.You can exit by POST/exit.")
		return message
	}
	if balance < 100 {
		fmt.Println("Balance less than 100. You can't withdraw.")
		message := fmt.Sprintf("Balance less than 100.You can't withdraw.You can exit by POST/exit")
		return message
	}
	checkValue, amount := checkInput(clientInput, balance)
	switch checkValue {
	case 0:
		fmt.Println("Error!Please enter ONLY ONE natural number")
		message := ("Error!Please enter ONE natural number")
		return message
	case 10:
		fmt.Println("Error!Please enter a Positive value")
		message := fmt.Sprintf("Error!Please enter a Positive value ")
		return message
	case 20:
		fmt.Println("Error!Please enter amount divisible by 100 ")
		message := fmt.Sprintf("Error!Please enter amount divisible by 100 ")
		return message
	case 30:
		fmt.Println("Error! Please enter amount less than or equal to 5000")
		message := fmt.Sprintf("Error! Please enter amount less than or equal to 5000 ")
		return message
	case 40:
		fmt.Printf("Error!Please enter amount less than or equal to your account balance that is: %v\n", balance)
		message := fmt.Sprintf("Error!Please enter amount less than or equal to your account balance")
		return message
	default:
		transactionCount++
		i := denominations(amount)
		balance = balance - amount
		fmt.Println("\n TRANSACTION SUCCESSFUL.YOUR BALANCE NOW IS\n:", balance)
		j := fmt.Sprintf("TRANSACTION SUCCESSFUL. ")
		message := i + j
		return message
	}
}
func checkInput(input int, balance int) (int, int) { //func to check valid amount given by client
	if input == 0 {
		return 0, input
	}
	if input < 0 {
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

func denominations(total int) string { //func to find out denominations of currency
	var q1 int
	var q2 int
	var q3 int
	q1 = total / 500
	total = total % 500
	q2 = total / 200
	total = total % 200
	q3 = total / 100
	fmt.Printf("500*%d +200*%d +100*%d \n", q1, q2, q3)
	i := fmt.Sprintf("500*%d +200*%d +100*%d ", q1, q2, q3)
	return i
}
