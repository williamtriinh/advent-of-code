package imath

func Max(x int, values ...int) int {
	max_int := x

	for _, value := range values {
		if value > max_int {
			max_int = value
		}
	}

	return max_int
}
