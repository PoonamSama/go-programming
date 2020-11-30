package main

import (
	"fmt"
)

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9} //THIS IS A SLICE OF INT
	x := sum(xi...)                     //WE ARE CALLING THE FUNC SUM HERE,PASSING A SLICE OF INT FOLLOWED BY ... WHICH CAN BE REFERRED TO UNFURLING A SLICE
	fmt.Println("The total is", x)
}

func sum(x ...int) int { //PARAMETER IS A NUMBER OF INTS(...int),NOT SLICE OF INT,so it wont accept xi
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}

// func (r receiver) identifier(parameter(s)) (return(s)) { code}
