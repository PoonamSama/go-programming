//Package mystr has catenate function
package mystr

import "strings"

//Catenate function will join separate words into a sentence
func Catenate(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {

		s += " "
		s += v

	}
	return s
}

//Join also joins the separate words
func Join(xs []string) string {
	return strings.Join(xs, " ")
}

//func Catenate and func Join,both do the same thing. we are going to compare their performances
