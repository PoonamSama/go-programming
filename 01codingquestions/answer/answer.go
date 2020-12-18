package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var balance int64 = 9563
	var s string
	count := 0
	s = "y"
	for s == "y" {

		x, amount := amt(balance)
		switch x {
		case 0:
			fmt.Println("Error!Please enter an integer")
			break

		case 10:
			fmt.Println("Error!Please enter a Positive value")
			break
		case 20:
			fmt.Println("Error!Please enter amount divisible by 100 ")
			break
		case 30:
			fmt.Printf("Error!Please enter amount less than or equal to your account balance that is: %v\n", balance)
			break
		case 40:
			fmt.Println("Please enter amount less than 5000")
			break
		default:
			count++
			//fmt.Println("count:", count)
			fmt.Println("You want to withdraw:", amount)
			fmt.Println("Your account balance:", balance)
			denoms(amount)
			balance = balance - amount
			fmt.Println("TRANSACTION SUCCESSFUL.YOUR BALANCE NOW IS:", balance)
			fmt.Println("------------------------")
			break
		}
		//fmt.Println(amount)
		//fmt.Println(count)
		if count >= 5 {
			e := fmt.Errorf("Error! you have made maximum number of transactions for today: 5")
			log.Fatalln(e)
		}
		if balance < 100 {
			e := fmt.Errorf("Error!your account balance is %v .It is less than 100.You can't withdraw", balance)
			log.Fatalln((e))
		}

		//	fmt.Println("value of s", s)
		s = foo()
		//fmt.Println("value of s", s)
		for s == "none" {
			s = foo()
		}
	}
	fmt.Println("Your total transactions today:", count)
}
func amt(y int64) (int64, int64) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the amount you want to withdraw:")
	scanner.Scan()

	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		return 0, input
	}
	if input < 0 {
		return 10, input
	}
	if input%100 != 0 {
		return 20, input
	}
	if y-input < 0 {
		return 30, input
	}
	if input > 5000 {
		return 40, input
	}
	return 50, input
}
func foo() string {
	var s string
	var x string

	fmt.Println("Do you want to make another transaction? enter y to continue and n to exit")
	fmt.Scanf("%s", &s)

	switch s {

	default:
		{
			x = "none"
			fmt.Println("Error!Wrong input. Enter y or n")
		}
	case "y":

		fmt.Println("You chose:yes")
		x = "y"

	case "n":

		fmt.Println("You chose:no; Exit")
		x = "n"

	}

	return x
}
func denoms(x int64) {
	var q1 int64
	var q2 int64
	var q3 int64
	q1 = x / 500
	x = x % 500
	q2 = x / 200
	x = x % 200
	q3 = x / 100
	fmt.Printf("500*%d +200*%d +100*%d \n", q1, q2, q3)
}
