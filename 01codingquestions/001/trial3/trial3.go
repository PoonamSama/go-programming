package main

import (
	"fmt"
	"log"
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
			e := fmt.Errorf("Error! you have made maximum transactions for today")
			log.Fatalln(e)
			//break
		}
		fmt.Println("Enter the amount you would like to withdraw")
		fmt.Scanf("%d", &amount)
		fmt.Println("Account Balance:", balance)
		if amount > 5000 {
			fmt.Println("ERROR! You can't withdraw more than 5000")
			count--
			continue
		}
		if amount%100 != 0 {
			fmt.Println("ERROR! Please enter a multiple of 100")
			count--
			continue
		}

		if balance < amount {
			fmt.Println("ERROR!Low balance!")
			fmt.Println("Your account balance is:", balance, "Please enter a valid amount.")
			count--
			continue
		}
		denoms(amount)
		balance = balance - amount
		fmt.Println("Transaction successful.Your balance now is:", balance)
		fmt.Println("-------------")

		fmt.Println("Please type yes if you wish to continue,and no to exit")
		fmt.Scanf("%s", &s)
		fmt.Println("you entered:", s)
		if s == "yes" {
			continue
		} else if s == "no" {
			fmt.Println("We are exiting")
			break
		} else {
			fmt.Println("Invalid command. Exit.")
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
