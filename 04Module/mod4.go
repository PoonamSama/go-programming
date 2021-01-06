package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type withdrawValue struct {
	AMOUNT int `json:"amount" `
}

func main() {
	balance := 9563
	transactionCount := 0
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Your balance is": getBalance(&balance)})
	})
	server.POST("/", func(c *gin.Context) {
		var clientInput withdrawValue
		c.BindJSON(&clientInput)
		c.JSON(200, gin.H{"Result": withdrawAmount(clientInput.AMOUNT, &balance, &transactionCount)}) // Your custom response here
	})
	server.POST("/exit", func(c *gin.Context) {
		os.Exit(1)
	})
	server.Run(":8080") // listen for incoming connections
}
func getBalance(balanceptr *int) string {
	message := fmt.Sprintf("%d", *balanceptr)
	return message
}
func withdrawAmount(clientVal int, balancePtr *int, counterPtr *int) string {
	if *counterPtr >= 5 {
		fmt.Println("Error! you have made maximum number of transactions for today: 5.")
		message := fmt.Sprintf("Maximum transactions reached for a day:5. You can exit using POST/exit.")
		return message
	}
	if *balancePtr < 100 {
		message := fmt.Sprintf("Balance less than 100.You can't withdraw anymore.You can exit using POST/exit.")
		return message
	}
	checkValue, amount := isAmountValid(clientVal, *balancePtr)
	switch checkValue {
	case 0:
		message := ("Error!Please enter ONE natural number")
		return message
	case 10:
		message := fmt.Sprintf("Error!Please enter a Positive value ")
		return message
	case 20:
		message := fmt.Sprintf("Error!Please enter amount divisible by 100 ")
		return message
	case 30:
		message := fmt.Sprintf("Error! Please enter amount less than or equal to 5000 ")
		return message
	case 40:
		message := fmt.Sprintf("Error!Please enter amount less than or equal to your account balance")
		return message
	default:
		*counterPtr++
		i := getDenominations(amount)
		*balancePtr = *balancePtr - amount
		j := fmt.Sprintf("TRANSACTION SUCCESSFUL.")
		message := i + j
		return message
	}
}
func isAmountValid(input int, balance int) (int, int) {
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

func getDenominations(total int) string {
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
