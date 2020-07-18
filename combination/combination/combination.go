//Package combination calculates the number of combinations(order does not matter) of n things taken r at a time
package combination

//Permuation takes n things taken r at a time and returns the result
func Combination(n int, r int) int {
	return factorial(n) / (factorial(n-r) * factorial(r))
}

//Factorial calculates the factorial of the given number
func factorial(n int) int {
	f := 1
	for ; n > 0; n-- {
		f *= n
	}
	return f
}
