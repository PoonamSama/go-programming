package main

import (
	"errors"
	"fmt"
	"log"
)

var errnewtype = errors.New("Error occurred,cant find root of negative number")

func main() {
	fmt.Printf("type of errnewtype %T:", errnewtype)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errnewtype
	}
	return 45, nil
}
