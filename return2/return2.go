package main

import (
	"fmt"
)

func main() {

	x := func() int { //anon func
		return 66
	}

	fmt.Printf("%T\n", x)
	y := x()
	fmt.Println(y)
	z := bar()
	fmt.Println(z) //cant print bcz its of type func() int
	fmt.Printf("%T\n", z)
	k := z() // k:=z wont work. z is of type func() int.so call that func by k:=z()
	fmt.Println(k)
	fmt.Printf("%T\n", k)

}

func bar() func() int { //return type of func bar is:func() int
	return func() int { //this func is an anon func
		return 451
	}
}
