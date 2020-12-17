package main

import (
	"fmt"
)

func main() {
	var amount int
	//var response string
	s := "y"
	for s == "y" {
		fmt.Println("enter amount")
		fmt.Scanf("%d", &amount)

		fmt.Println(amount)
		s = foo()
		for s == "none" {
			s = foo()
		}

		/*	fmt.Println("enter y or n")
			fmt.Scanf("%s", &s)
			switch s {
			case "y":

				fmt.Println("yes")

			case "n":

				fmt.Println("no")
				break

			default:
				fmt.Println("wrong choice")
				fmt.Println("enter y or n")
				fmt.Scanf("%s", &s)

			}*/

	}
}
func foo() string {
	var s string
	var x string

	fmt.Println("enter y or n")
	fmt.Scanf("%s", &s)
	/*	if v2 == 0 {
		x = "none"
		return x
	}*/
	switch s {

	default:
		{
			x = "none"
			fmt.Println("wrong input")
		}
	case "y":

		fmt.Println("yes")
		x = "y"

	case "n":

		fmt.Println("no")
		x = "n"

		//x = "none"
		//fmt.Println("wrong choice")

		//fmt.Println("enter y or n")
		//	fmt.Scanf("%s", &s)
		//foo()

	}
	//fmt.Println(x)
	return x
}
