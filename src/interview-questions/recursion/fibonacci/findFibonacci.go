package fibonacci

func findFibonacciNumber(iteration int) int {
	if iteration <= 1 {
		return iteration
	}
	return findFibonacciNumber(iteration-1) + findFibonacciNumber(iteration-2)
}
