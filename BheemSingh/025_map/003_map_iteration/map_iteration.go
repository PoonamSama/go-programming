package main

import "fmt"

func main() {

	map_1 := map[int]string{

		90: "banana",
		91: "apple",
		92: "mango",
		93: "grapes",
		94: "chikoo",
	}

	for fruit, name := range map_1 {
		fmt.Println(fruit, name)
	}

}
