//Package acdc adds a number of ints
package acdc

//Sum adds ints xi ...int and returns int
func Sum(xi ...int) int {
	x := 0
	for _, v := range xi {
		x += v
	}
	return x
}
