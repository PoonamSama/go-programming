package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var balance = 9563
var count = 0

//LOGIN exported
type LOGIN struct {
	AMOUNT int `json:"amount"`
}

func main() {
	r := gin.Default()
	r.GET("/check", func(c *gin.Context) {
		c.String(200, "Welcome")
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Your balance is": balance})
	})
	r.POST("/", func(c *gin.Context) {
		var login LOGIN
		c.BindJSON(&login)
		c.JSON(200, gin.H{"Result": postresponds(login.AMOUNT)})
	})
	r.POST("/exit", func(c *gin.Context) {
		c.JSON(200, "Exit")
		os.Exit(1)
	})
	r.Run(":8080") // listen for incoming connections
}
func postresponds(amountInt int) string { //response func for method POST/
	if count >= 5 {
		fmt.Println("Error! you have made maximum number of transactions for today: 5.")
		errj := fmt.Sprintf("Maximum transactions reached for a day:5.You can exit by POST/exit.")
		return errj
	}
	if balance < 100 {
		fmt.Println("Balance less than 100. You can't withdraw.")
		errj := fmt.Sprintf("Balance less than 100.You can't withdraw.You can exit by POST/exit")
		return errj
	}
	x, amount := amt(amountInt, balance)
	switch x {
	case 0:
		fmt.Println("Error!Please enter ONLY ONE natural number")
		errj := ("Error!Please enter ONE natural number")
		return errj
	case 10:
		fmt.Println("Error!Please enter a Positive value")
		errj := fmt.Sprintf("Error!Please enter a Positive value ")
		return errj
	case 20:
		fmt.Println("Error!Please enter amount divisible by 100 ")
		errj := fmt.Sprintf("Error!Please enter amount divisible by 100 ")
		return errj
	case 30:
		fmt.Println("Error! Please enter amount less than or equal to 5000")
		errj := fmt.Sprintf("Error! Please enter amount less than or equal to 5000 ")
		return errj
	case 40:
		fmt.Printf("Error!Please enter amount less than or equal to your account balance that is: %v\n", balance)
		errj := fmt.Sprintf("Error!Please enter amount less than or equal to your account balance")
		return errj
	default:
		count++
		i := denoms(amount)
		balance = balance - amount
		fmt.Println("\n TRANSACTION SUCCESSFUL.YOUR BALANCE NOW IS\n:", balance)
		j := fmt.Sprintf("TRANSACTION SUCCESSFUL. ")
		errj := i + j
		return errj
	}
}
func amt(input int, balance int) (int, int) { //func to check valid amount given by client
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

func denoms(x int) string { //func to find out denominations of currency
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
