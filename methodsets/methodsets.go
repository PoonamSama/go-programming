//create a person struct
//create a func called “changeMe” with a *person as a parameter
//change the value stored at the *person address
package main

import (
	"fmt"
)

type person struct {
	name string
}

func changeme(x *person) {
	x.name = "James"
}
func main() {
	p1 := person{
		name: "Poonam",
	}
	fmt.Println(p1)
	changeme(&p1)
	fmt.Println(p1)
}
