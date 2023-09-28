package find_max_consecutive_ones

func findMaxConsecutiveOnes(nums []int) int {
	var (
		count    int
		maxCount int
	)
	for i := range nums {
		if 1 == nums[i] {
			count = count + 1
		} else {
			if maxCount < count {
				maxCount = count
			}
			count = 0
		}
	}
	if maxCount < count {
		maxCount = count
	}
	return maxCount
}
