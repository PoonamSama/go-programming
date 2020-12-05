package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string //if we write like first (lowercase),o/p would be empty[{}{}]
	Last  string //can only access if it starts with capital letter
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
