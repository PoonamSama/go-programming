package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}
func sqrt(x float64) (float64, error) {
	y := fmt.Errorf("An error has occurred: cant use negative number %v", x)
	if x < 0 {
		return 0, y
	}
	return 45, nil
}
