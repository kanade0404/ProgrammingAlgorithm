package first_letter_to_appear_twice

func repeatedCharacter(s string) byte {
	m := make(map[byte]int)
	var result byte
	for i := range s {
		if _, ok := m[s[i]]; ok {
			result = s[i]
			break
		} else {
			m[s[i]] = 1
		}
	}
	return result
}
