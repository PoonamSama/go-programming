//you can't have blank lines
//otherwise they wont be printed

//Package mymath provides ACME math solutions :must start with Package <pkgname>
package mymath

//Sum adds a number of integers and returns an int:must start with a capital letter function name;cap means can be exported
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
