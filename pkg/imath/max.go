package imath

func Max(values ...int) int {
	max := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}

	return max
}
