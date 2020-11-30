package main

import (
	"fmt"
)

type person struct { //keyword identifier typeofvariable
	first string
	last  string
}

// var x int : keyword identifier type(int,struct,interface,bool etc)
type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "from secretAgent speak")
}
func (s person) speak() {
	fmt.Println("I am", s.first, s.last, "from person speak")
}

type human interface { //now the values in sa1{} are of type human too
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into func bar", h)
}
func main() {
	sa1 := secretAgent{ //the value in{} is of type secretAgent;and a value can be of more than one type;sa1 is variable
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	p1 := person{
		"Dr.",
		"Yes",
	}

	fmt.Println(sa1)
	//sa1.speak()
	//sa2.speak()
	//p1.speak()
	bar(sa1) //the fn bar can take type secretAgent and type person : this is called polymorphism
	bar(sa2) //because the fn bar takes type human
	bar(p1)  // and type human interface takes speak(); and speak (r receives) both secretAgent and person
}
