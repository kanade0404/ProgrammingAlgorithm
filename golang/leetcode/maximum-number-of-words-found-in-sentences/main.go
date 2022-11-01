package maximum_number_of_words_found_in_sentences

import "strings"

func mostWordsFound(sentences []string) int {
	var result int
	for _, sentence := range sentences {
		c := strings.Count(sentence, " ") + 1
		if c > result {
			result = c
		}
	}
	return result
}
