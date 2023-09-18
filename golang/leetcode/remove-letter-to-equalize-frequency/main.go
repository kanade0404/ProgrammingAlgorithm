package remove_letter_to_equalize_frequency

import (
	"strings"
)

func equalFrequency(word string) bool {
	strs := strings.Split(word, "")
	m := make(map[string]int)
	for _, str := range strs {
		if count, ok := m[str]; ok {
			m[str] = count + 1
		} else {
			m[str] = 1
		}
	}
	for s1, _ := range m {
		cm := m
		cm[s1] = 0
		var freq int
		for _, i2 := range cm {
			if i2 == 0 {
				continue
			}
			if freq == 0 {
				freq = i2
				continue
			} else {
				if freq == i2 {

				}
			}
		}
	}
	return false
}
