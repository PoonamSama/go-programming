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
	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}
	for _, v := range m {
		fmt.Println(v.first_name)
		fmt.Println(v.last_name)
		for i, value := range v.fav_ice {
			fmt.Println(i, value)

		}
	}
}
