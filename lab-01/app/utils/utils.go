package utils

func MinOfInt(values ...int) int {
	min := values[0]

	for _, i := range values {
		if min > i {
			min = i
		}
	}
	return min
}

func Max2Int(v1, v2 int) int {
    if v1 < v2 {
        return v2
    }
    return v1
}
