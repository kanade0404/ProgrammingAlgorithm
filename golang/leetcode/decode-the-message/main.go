package decode_the_message

func decodeMessage(key string, message string) string {
	var (
		count  int
		result = make([]rune, 0, len(message))
		m      = make(map[byte]int)
	)
	for i := range key {
		if key[i] == 32 {
			continue
		}
		if _, ok := m[key[i]]; !ok {
			m[key[i]] = count + 97
			count++
		}
	}
	for i := range message {
		if message[i] == 32 {
			result = append(result, ' ')
		} else {
			result = append(result, rune(m[message[i]]))
		}
	}
	return string(result)
}
