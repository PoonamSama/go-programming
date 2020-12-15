package main

import (
	"fmt"
	"strings"

	mystr "github.com/GoesToEleven/go-programming/Benchmarking/002cat/pkgmystr"
)

const s = "A constant string that has many words,therefore creating a sentence. The sentence in the present tense,may not make any sense but don't be so tense"

func main() {
	xs := strings.Split(s, " ")
	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n %s \n", mystr.Catenate(xs))
	fmt.Printf("\n %s \n", mystr.Join(xs))
}
