package main

import (
	"fmt"

	dog "github.com/GoesToEleven/go-programming/NinjaLevel13/ex1/pkgdog"
)

type canine struct {
	name string
	age  int
}

func main() {
	d1 := canine{
		"Husky",
		dog.Years(10),
	}
	fmt.Println(d1)
	d2 := canine{
		"Chihuahua",
		dog.Years(7),
	}
	fmt.Println(d2)

}
