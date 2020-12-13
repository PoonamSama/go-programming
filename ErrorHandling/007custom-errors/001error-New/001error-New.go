package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Error:cant find sqrt of negative num")
	}
	return 42, nil
}
