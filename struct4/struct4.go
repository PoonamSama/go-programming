package main

import (
	"fmt"
)

type person struct {
	first_name string
	last_name  string
	fav_ice    []string
}

func main() {
	p1 := person{
		first_name: "Poonam",
		last_name:  "Gupta",
		fav_ice:    []string{"vanilla", "choco"},
	}
	fmt.Println(p1)
	p2 := person{
		first_name: "James",
		last_name:  "Bond",
		fav_ice:    []string{"Strawberry", "Mint"},
	}
	fmt.Println(p2)
	for i, v := range p2.fav_ice {
		fmt.Println(i, v)
	}
}
