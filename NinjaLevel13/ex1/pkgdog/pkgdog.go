//Package dog has two functions that convert human yrs into dog yrs
package dog

//Years this func simply converts human yrs into dog yrs
func Years(x int) int {
	return x * 7
}

//YearsTwo uses loop to multiply
func YearsTwo(y int) int {
	count := 0
	for i := 0; i < y; i++ {
		count += 7
	}
	return count
}
