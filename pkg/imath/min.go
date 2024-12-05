package imath

func Min(values ...int) int {
	min := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}

	return min
}
