package shuffle_string

func restoreString(s string, indices []int) string {
	strLen := len(s)
	results := make([]rune, strLen, strLen)
	ss := []rune(s)
	for i := 0; i < strLen; i++ {
		results[indices[i]] = ss[i]
	}
	return string(results)
}
