package main

import (
	"fmt"
)

func main() {
	var balance int = 9563
	var amount int

	var s string
	count := 0
	s = "yes"
	for s == "yes" {
		count++
		fmt.Println("counter:", count)
		if count > 5 {
			fmt.Println("ERROR! You have made maximum transactions for today.")
			break
		}
		fmt.Println("Enter the amount you would like to withdraw")
		fmt.Scanf("%d", &amount)
		fmt.Println("Account Balance:", balance)
		if amount > 5000 {
			fmt.Println("ERROR! You can't withdraw more than 5000")
			break
		}
		if amount%100 != 0 {
			fmt.Println("ERROR! Please enter a multiple of 100")
			break
		}

		if balance < amount {
			fmt.Println("ERROR!Low balance!")
			break
		}
		denoms(amount)
		balance = balance - amount
		fmt.Println("Transaction successful.Your balance now is:", balance)

		fmt.Println("Please type yes if you wish to continue")
		fmt.Scanf("%s", &s)
		fmt.Println("you entered:", s)
		if s == "yes" {
			continue
		} else {
			fmt.Println("we are exiting")
			break
		}

	}
}
func denoms(x int) {
	var q1 int
	var q2 int
	var q3 int
	q1 = x / 500
	x = x % 500
	q2 = x / 200
	x = x % 200
	q3 = x / 100
	fmt.Printf("500*%d +200*%d +100*%d \n", q1, q2, q3)
}
