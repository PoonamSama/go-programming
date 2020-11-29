package main

import (
	"fmt"
)

func main() {
	const name, dept = "Poonam Gupta", "INS"
	s := fmt.Sprintf("The name is %s and department is %s", name, dept)
	fmt.Println(s)
}
