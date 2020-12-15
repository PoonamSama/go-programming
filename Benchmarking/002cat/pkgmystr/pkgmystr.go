//Package mystr has catenate function
package mystr

import "strings"

//Catenate function will join separate words into a sentence
func Catenate(xs []string) string {
	s := ""
	for _, v := range xs {
		s += v
		s += " "
	}
	return s
}
func Join(xs []string) string {
	return strings.Join(xs, " ")
}

//func Catenate and func Join,both do the same thing. we are going to compare their performances
