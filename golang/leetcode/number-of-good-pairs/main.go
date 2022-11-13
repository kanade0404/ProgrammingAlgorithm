package number_of_good_pairs

func numIdenticalPairs(nums []int) int {
	var result int
	for i := range nums {
		for j := range nums {
			if nums[i] == nums[j] && i < j {
				result += 1
			}
		}
	}
	return result
}
