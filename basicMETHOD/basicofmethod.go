package main

import (
	"fmt"
)

type female struct {
	name string
	age  int
}
type male struct {
	firstname string
	old       int
}
type team struct {
	female
	male
	ltk bool
}

func (f female) talk() {
	fmt.Println("I am a woman", f.name, f.age) //f.fieldname that's how the method accesses struct it receives
}
func (f male) talk() {
	fmt.Println("I am a man", f.firstname, f.old)
}
func (f team) talk() {
	fmt.Println("We are a team", f.name, f.age, f.firstname, f.old, f.ltk)
}

type people interface {
	talk()
}

func bar(h people) {
	h.talk()
}
func main() {

	fmt.Println(" h")
	f1 := female{
		name: "Poonam Gupta",
		age:  25,
	}
	m1 := male{
		firstname: "James Bond",
		old:       32,
	}
	t1 := team{
		female: female{
			name: "jia",
			age:  21,
		},
		male: male{
			firstname: "Lawliet",
			old:       23,
		},
		ltk: true,
	}
	bar(f1)
	bar(m1)
	bar(t1)
	t1.talk() //value t1 has access to method talk: basic of method
}
