package main

import (
	"fmt"
)

type people struct {
	firstname string
	lastname  string
	age       int
}
type secretAgent struct {
	people //we don't name here
	ltk    bool
}

func main() {
	sa1 := secretAgent{
		people: people{ //but we have to name here, to assign it some value
			firstname: "James",
			lastname:  "Bond",
			age:       33,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa1.age)                      //this is easier;age got promoted
	fmt.Println("Kill", sa1.people.firstname) //this is also correct

}
