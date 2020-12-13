package main

import (
	"fmt"
	"log"
)

/*type error interface{
	Error() string
}*/
type mynewerr struct { //if we hv a type which has the method Error(),
	//then any value of that type implicitly implements the error interface
	name  string
	place string
	err   error
}

func (m mynewerr) Error() string { //thismethod receives type mynewerr and returns a string.Error()
	//method Error() implements the interface
	return fmt.Sprintf("error occurred: %v %v %v", m.name, m.place, m.err)
}
func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}
func sqrt(x float64) (float64, error) {

	if x < 0 {

		z := fmt.Errorf("We encountered error for square root of negative number %v", x) //this is the err error in mynewerr struct
		return 0, mynewerr{
			"Jessi",
			"Animeland",
			z,
		}

	}
	return 455, nil
}
