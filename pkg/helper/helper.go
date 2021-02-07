package helper

// RemoveString returns []string with the given `r` removed from it
func RemoveString(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}

	return s
}

// Reverse returns a reversed []string
func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// Sum returns the sum of []int
func Sum(array []int) int {
	result := 0

	for _, v := range array {
		result += v
	}

	return result
}
