package utils

func IntSliceSum(value []int) (result int) {
	for _, r := range value {
		result += r
	}
	return
}

