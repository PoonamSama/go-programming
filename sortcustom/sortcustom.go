package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}
type ByAge []person                //byage is a slice of type person
func (a ByAge) Len() int           { return len(a) }              // these methods are attached to type byage
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }    //we copied this from example of
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age } //package sort
func main() {
	p1 := person{"Poonam", 25}
	p2 := person{"James", 11}
	p3 := person{"Jessi", 33}
	people := []person{p1, p2, p3}
	fmt.Println(people)
	sort.Sort(ByAge(people)) //converting people into byAge so that people has methods attached to it
	fmt.Println(people)
}
