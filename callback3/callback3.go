package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("yes it worked!")
	}
	y := func() {
		fmt.Println("no,please try again")
	}
	thisorthat(x, y) //calling the func and passing the other funcs
}
func thisorthat(yes func(), no func()) { //yes is the name of func passed to thisorthat
	worked := true
	if worked == false {
		yes()
	} else {
		no()
	}
}
