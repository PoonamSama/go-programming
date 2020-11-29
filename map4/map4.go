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
	m["Todd"] = 35 //adding "Todd" to map m
	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println(m)
	delete(m, "jb") //TO DELETE A KEY FROM A MAP
	fmt.Println(m)
}
