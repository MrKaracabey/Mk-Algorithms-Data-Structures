package multiplyNumbers

func productNumbers(number1, number2 int) int {
	if number2 == 0 {
		return 0
	}

	return productNumbers(number1, number2-1) + number1
}
