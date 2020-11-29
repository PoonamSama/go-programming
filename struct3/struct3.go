package main

import (
	"fmt"
)

func main() {
	p1 := struct { //Anonymous struct;has no name
		firstname string
		lastname  string
		age       int
	}{
		firstname: "James",
		lastname:  "Bond",
		age:       21,
	}
	fmt.Println(p1)
}
