package two_sum

func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				if i < j {
					result = []int{i, j}
				} else {
					result = []int{j, i}
				}
				break
			}
		}
	}
	return result
}
