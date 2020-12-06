package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 34, 56, 11, 2, 0, 9}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	xs := []string{"James Bond", "MoneyPenny", "Dr. Yes", "Q", "q"}
	sort.Strings(xs)
	fmt.Println(xs)

}
