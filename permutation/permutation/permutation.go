//Package permutation calculates the number of permutations(order matters) of n things taken at a time
package permutation

//Permuation takes n things taken r at a time and returns the result
func Permutation(n int, r int) int {
	return factorial(n) / factorial(n-r)
}

//Factorial calculates the factorial of the given number
func factorial(n int) int {
	f := 1
	for ; n > 0; n-- {
		f *= n
	}
	return f
}
