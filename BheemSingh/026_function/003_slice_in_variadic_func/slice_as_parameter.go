package main

import "fmt"

func main() {

	x := []string{"Go", "is", "all", "about", "types"}

	str(x...)

}

func str(s ...string) {
	fmt.Println(s)
}
