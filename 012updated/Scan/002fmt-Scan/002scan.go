package main

import "fmt"

func main() {
	var xs [7]string
	fmt.Println("Start the Army fanchant!")
	/*for i := 0; i < 7; i++ {
		fmt.Println("Name:")
		fmt.Scanf("%s ", &xs[i])
	}
	fmt.Println("Here we go!")
	for i := 0; i < 7; i++ {

		fmt.Println(xs[i])
	}
	fmt.Println("BTS!!!")*/
	xs[3] = "Kim Namjoon"
	fmt.Print(xs)
}
