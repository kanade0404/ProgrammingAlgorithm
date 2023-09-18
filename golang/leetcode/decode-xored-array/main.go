package decode_xored_array

func decode(encoded []int, first int) []int {
	var results []int
	results = append(results, first)
	before := first
	for i := range encoded {
		e := encoded[i]
		r := before | e
		results = append(results, r)
		before = r
	}
	return results
}
