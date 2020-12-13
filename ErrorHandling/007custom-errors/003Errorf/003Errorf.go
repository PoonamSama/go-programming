package main

import (
	"fmt"
	"log"
	//"errors"
)

func main() {

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("Error occurred,cant find root of negative number %v", x)
	}
	return 45, nil
}
