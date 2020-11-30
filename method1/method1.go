package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	ltk bool
}

//func(r receiver) identifier(parameters)(return(s)){code}
func (s secretAgent) speak() { //receiver s has access to all of secretAgent and all data that's of type secretAgent
	fmt.Println("I am", s.first, s.last)
	fmt.Println("And I can kill that's", s.ltk)
}

func main() {
	p1 := secretAgent{
		person: person{first: "James",
			last: "Bond",
			age:  32,
		},
		ltk: true,
	}
	p2 := secretAgent{
		person: person{first: "Poonam",
			last: "Gupta",
			age:  25,
		},
		ltk: true,
	}
	fmt.Println(p1)
	p1.speak() //giving func speak the data of p1
	p2.speak()
}
