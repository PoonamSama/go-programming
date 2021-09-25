package main

import (
	"fmt"
)

func main() {
	var day int

	fmt.Println("Enter day as integer")

	fmt.Scanf("%d", &day)

	switch {
	case day == 1:
		fmt.Println("Monday")
	case day == 2:
		fmt.Println("Tuesday")
	case day == 3:
		fmt.Println("Wednesday")
	case day == 4:
		fmt.Println("Thursday")
	case day == 5:
		fmt.Println("Friday")
	case day == 6:
		fmt.Println("Saturday")
	case day == 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}
}
