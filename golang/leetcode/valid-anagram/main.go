package valid_anagram

import "sort"

func isAnagram(s string, t string) bool {
	bs, bt := []byte(s), []byte(t)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	sort.Slice(bt, func(i, j int) bool {
		return bt[i] < bt[j]
	})
	return string(bs) == string(bt)
}
