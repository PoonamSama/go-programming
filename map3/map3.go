package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"jb": 32,
		"mp": 25,
		"df": 78,
	}
	fmt.Println(m)
	fmt.Println(m["jb"])
	fmt.Println(m["df"])
	fmt.Println(m["xy"]) //key wasnt in the map still gave a value:0 this isn't good
	fmt.Println(m)
	//to check if a key is in map or not:ok(bool )
	v, ok := m["xy"] // v stores the value of key "xy"
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["jb"]; ok {
		fmt.Println("This is the if print", v)
	}

}
