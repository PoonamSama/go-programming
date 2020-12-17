package main

import (
	"fmt"
	"log"
)

var totaltransaction int
var count = 0

func main() {
	var balance int = 9563
	var amount int

	var s string

	s = "yes"
	for s == "yes" {
		count++
		//fmt.Println("counter:", count)

		if count > 5 {
			count--
			totaltransaction = count
			fmt.Println("Your total transactions for today:", totaltransaction)
			e := fmt.Errorf("Error! you have made maximum transactions for today")
			log.Fatalln(e)

		}
		fmt.Println("Enter the amount you would like to withdraw")
		fmt.Scanf("%d", &amount)
		fmt.Println("Account Balance:", balance)
		if amount > 5000 {
			count--
			totaltransaction = count
			fmt.Println("Your total transactions for today:", totaltransaction)
			err := fmt.Errorf("ERROR! You can't withdraw more than 5000")
			log.Fatalln(err)

		}
		if amount%100 != 0 {
			count--
			totaltransaction = count
			fmt.Println("Your total transactions for today:", totaltransaction)
			err := fmt.Errorf("ERROR! Please enter a multiple of 100")
			log.Fatalln(err)

		}

		if balance < amount {
			count--
			totaltransaction = count
			fmt.Println("Your total transactions for today:", totaltransaction)
			err := fmt.Errorf("ERROR !Low balance")
			fmt.Println("Your account balance is:", balance, "Please enter a valid amount.")
			log.Fatalln(err)

		}
		denoms(amount)
		balance = balance - amount
		fmt.Println("Transaction successful.Your balance now is:", balance)
		totaltransaction = count
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
	fmt.Println("your Total Transactions for today are:", totaltransaction)
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

// this code doesn't count lowbalance,max amount of transaction,not100multiples, as number of
//transactions,but it exits once an invalid value is input
