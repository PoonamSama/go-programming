package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println(x)
	foo(&x)
	fmt.Println(x)

}
func foo(y *int) {  //THIS IS PASS BY VALUE
	fmt.Println(*y)
	*y = 45
	fmt.Println(*y)

}
