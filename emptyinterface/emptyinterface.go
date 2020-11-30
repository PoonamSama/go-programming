package main

import (
	"fmt"
)

type person struct { //keyword identifier typeofvariable
	first string
	last  string
}
type poonam struct { //keyword identifier typeofvariable
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func (s secretAgent) speak() { //speak is the METHOD; secretAgent has METHOD speak()
	fmt.Println("I am", s.first, s.last, "from secretAgent speak")
}
func (s person) speak() { //person also has has METHOD speak()
	fmt.Println("I am", s.first, s.last, "from person speak")
}

type human interface { //an empty interface, w/o any method

}

func bar(h human) {
	fmt.Println("I was passed into func bar", h)
	switch h.(type) {
	case secretAgent:
		fmt.Println("I was passed in barrrr", h.(secretAgent).last) //h.(secretAgent is assertion, not conversion like in type hotdog int)
	case person:
		fmt.Println("I was passed in barrrr", h.(person).last)
	}
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
	p2 := poonam{
		"lady",
		"baby",
	}
	fmt.Println(sa1)
	//sa1.speak()
	//sa2.speak()
	//p1.speak()
	bar(sa1) //the fn bar can take type secretAgent and type person : this is called polymorphism
	bar(sa2) //because the fn bar takes type human
	bar(p1)  // and type human interface takes speak(); and speak (r receives) both secretAgent and person
	bar(p2)  // didn't have any method ,but passed because human is an empty interface
}
