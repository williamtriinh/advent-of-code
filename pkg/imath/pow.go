package imath

func Pow(base, exponent int) int {
	if exponent == 0 {
		return 1
	}

	if exponent == 1 {
		return base
	}

	result := base
	for i := 2; i <= exponent; i++ {
		result *= base
	}

	return result
}
