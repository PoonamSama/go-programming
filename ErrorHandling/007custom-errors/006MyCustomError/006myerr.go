package main

import (
	"fmt"
	"log"
)

var choice string

type person struct {
	name string
	age  int
	err  error
}

func (c person) Error() string {
	return fmt.Sprintf("ERROR! %v %v %v ", c.name, c.age, c.err)
}
func main() {

	fmt.Println("Enter your choice:")
	fmt.Scan(&choice)
	_, err := thisfunc(choice)
	if err != nil {
		log.Fatalln(err)
	}
}
func thisfunc(s string) (int, error) {
	z := fmt.Errorf("You entered wrong choice %v", s)
	if s == "male" {
		return 0, person{
			"Not Acceptable",
			26,
			z,
		}
	} else if s == "female" {
		return 0, person{
			"unacceptable",
			222,
			z,
		}

	}
	return 999, nil
}
