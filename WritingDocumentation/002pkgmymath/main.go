// package mynewwmathtwo provides ACME
package mynewwmathtwo

// i commented this line
func Sum(xi ...int) int {
	sum := 0
	for v := range xi {
		sum += v
	}
	return sum
}
