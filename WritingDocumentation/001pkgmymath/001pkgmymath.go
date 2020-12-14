//you can't have blank lines
//otherwise they wont be printed

//Package mymath provides ACME math solutions
package mymath

//Sum adds a number of integers and returns an int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
