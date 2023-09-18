package group_anagrams

func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string]map[uint8]int)
	for i := range strs {
		if _, ok := resultMap[strs[i]]; !ok {
			resultMap[strs[i]] = make(map[uint8]int)
		}
		for c := range strs[i] {
			s := strs[i][c]
			if _, ok := resultMap[strs[i]][s]; ok {
				resultMap[strs[i]][s] += 1
			} else {
				resultMap[strs[i]][s] = 1
			}
		}
	}
	for key := range resultMap {

	}
	return [][]string{}
}
