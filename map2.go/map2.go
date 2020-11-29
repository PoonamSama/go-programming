package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Miss":  45,
		"Lia":   21,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["queen"])
	v3, ok3 := m["queen"] //bool ok3 is false
	fmt.Println(v3, ok3)
	if v, ok := m["Kia"]; ok {
		fmt.Println("Kia is a key of value:", v) //this wont print since bool ok is false
	}
	if v1, ok1 := m["Lia"]; ok1 {
		fmt.Println("Lia is a key of value:", v1)
	}

}
