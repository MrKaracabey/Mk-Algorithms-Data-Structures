package factorial

func findFactorial(number int) int {

	if number == 0 {
		return 1
	}

	return findFactorial(number-1) * number
}
