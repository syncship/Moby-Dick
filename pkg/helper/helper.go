package helper

// RemoveString returns []string with the given `r` removed from it
// TODO: remove comment, isto tava alterando o slice original s []string
// pode ser a intençao mas geralmente nao é
func RemoveString(s []string, r string) []string {
	ret := make([]string, 0, len(s))
	for _, v := range s {
		if v == r {
			continue
		}
		ret = append(ret, v)
	}

	return ret
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
