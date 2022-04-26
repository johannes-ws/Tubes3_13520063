package BM

import "math"

func buildLast(pattern string) []int {
	var last [128]int
	for i := 0; i < 128; i++ {
		last[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		last[int(pattern[i])] = i
	}
	return last[:]
}

func bmMatch(text string, pattern string) int {
	last := buildLast(pattern)
	n := len(text)
	m := len(pattern)
	i := m - 1
	if i > n-1 {
		return -1
	}
	j := m - 1
	for {
		if pattern[j] == text[i] {
			if j == 0 {
				return i
			} else {
				i--
				j--
			}
		} else {
			lo := last[int(text[i])]
			i += m - int(math.Min(float64(j), float64(1+lo)))
			j = m - 1
		}
		if i > n-1 {
			break
		}
	}
	return -1
}
