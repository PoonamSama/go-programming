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

	fmt.Println(map_1)

	delete(map_1, 92)

	fmt.Println(map_1)

}
