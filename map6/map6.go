package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		"James":  {"hi", "bye", "hello"},
		"Jessi":  {"Team", "Rocket", "True"},
		"Meowth": {"Leader", "Talks", "Pokemon"},
	}
	for key, value := range x {

		fmt.Println("For the key value", key)
		for index, v := range value {
			fmt.Println(index, v)
		}
	}
	x["Ash"] = []string{"Pika", "chu"}
	for k, v3 := range x {
		fmt.Println(k, v3)
	}
}
