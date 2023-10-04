package merge_strings_alternately

import "math"

func mergeAlternately(word1 string, word2 string) string {
	l := int(math.Max(float64(len(word1)), float64(len(word2))))
	var result []byte
	for i := 0; i < l; i++ {
		if i < len(word1) {
			result = append(result, word1[i])
		}
		if i < len(word2) {
			result = append(result, word2[i])
		}
	}
	return string(result)
}
