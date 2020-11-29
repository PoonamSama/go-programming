package main

import (
	"fmt"
)

type people struct { //composite or aggregate data type
	firstname string
	lastname  string
	age       int
}

func main() {
	p1 := people{
		firstname: "James",
		lastname:  "Bond",
		age:       33,
	}
	p2 := people{
		firstname: "Poonam",
		lastname:  "Gupta",
	}
	p3 := people{}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.age)
	fmt.Println(p2.firstname)
	fmt.Println(p3)
}
