package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
	ltk  bool
}

func main() {
	p1 := person{"Poonam", 26, false}
	fmt.Println(p1.age, p1.name, p1.ltk)
	fmt.Println(p1)

}
