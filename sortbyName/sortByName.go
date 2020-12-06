package main
import (
	"fmt"
	"sort"
)
type person struct{
	name string
	age int
}
type ByName []person
func (ab ByName) Len() int           { return len(ab) }              // these methods are attached to type byage
func (ab ByName) Swap(i, j int)      { ab[i], ab[j] = ab[j], ab[i] }    //we copied this from example of
func (ab ByName) Less(i, j int) bool { return ab[i].name < ab[j].name } //package sort
func main(){
	p1:=person{"Poonam",25}
	p2:=person{"James",111}
	p3:=person{"Harriet",32}
	people:=[]person{p1,p2,p3} 
	sort.Sort(ByName(people))
	fmt.Println(people)
}