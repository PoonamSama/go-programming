package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{ //key is string and value is slice of strings
		"name":   {"likes", "comments", "shares"},
		"Poonam": {"success", "inspiration", "dreams"},
		"Oscar":  {"wins", "satisfaction", "happy"},
	}

	fmt.Println(x)
	for key, value := range x {
		fmt.Println("For the key value", key)
		for i, v := range value { //to access the slice of strings
			fmt.Println(i, v)
		}
	}
}
