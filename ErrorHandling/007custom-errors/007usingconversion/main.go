package main

import (
	"fmt"
)

type customer struct {
	name    string
	order   string
	mistake string
}

func (c customer) Error() string {
	return fmt.Sprintf("Error occurred! Customer name: %v Customer Order: %v Your Mistake: %v", c.name, c.order, c.mistake)
} //when commented out,code doesnt run bcz conversion not possible
func main() {
	c1 := customer{
		"Poonam Gupta",
		"Frappucchino",
		"Coffee",
	}
	foo(c1)

}
func foo(e error) {
	fmt.Println(e) //conversion is happening here, since inherent type is error for e error and also c1 (type customer is an error because it implements that method and interface)
}
