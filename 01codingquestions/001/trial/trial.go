package main

import "fmt"

func main() {
	var s string
	fmt.Println("Please enter a choice: yes or no")
	fmt.Scanf("%s", &s)
	fmt.Println("you entered:", s)
	for s == "yes" {
		fmt.Println("continuing the transaction")
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
