package split_a_string_in_balanced_strings

func balancedStringSplit(s string) int {
	var (
		rLen, lLen, count int
	)
	for i := range s {
		if s[i] == 'R' {
			rLen++
		} else {
			lLen++
		}
		if rLen == lLen {
			count++
			rLen = 0
			lLen = 0
		}
	}
	return count
}
