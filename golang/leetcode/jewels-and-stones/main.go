package jewels_and_stones

func numJewelsInStones(jewels string, stones string) int {
	jewelsLen := len(jewels)
	stonesLen := len(stones)
	var result int
	for i := 0; i < stonesLen; i++ {
		for j := 0; j < jewelsLen; j++ {
			if stones[i] == jewels[j] {
				result += 1
				break
			}
		}
	}
	return result
}
