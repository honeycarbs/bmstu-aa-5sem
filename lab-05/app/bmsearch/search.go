package bmsearch

func ConstructSkipTable(pattern []rune) map[rune]int {
	length := len(pattern)
	skips := make(map[rune]int)
	for i, char := range pattern {
		if i < length - 1 {
			skips[char] = length - i - 1
		} else {
			skips[char] = length
		}
	}
	return skips
}

func FindFirstIndex(source, pattern []rune, skips map[rune]int) int {
	len_ptrn := len(pattern)
	len_src := len(source)

	j := len_ptrn - 1
	for j < len_src {
		k := j
		i := len_ptrn - 1
		for i >= 0 {
			if source[k] != pattern[i] {
				if skip, found := skips[source[j]]; found {
					j += skip
				} else {
					j += len_ptrn
				}
				break
			}
			k--
			i--
		}
		if i < 0 {
			return k + 1
		}
	}
	return -1
}
