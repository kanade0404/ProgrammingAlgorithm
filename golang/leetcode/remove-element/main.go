package remove_element

func removeElement(nums []int, val int) int {
	var result int
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		if nums[i] == val {
			nums[i] = -1
		} else {
			result += 1
		}
	}
	var cpNums []int
	for i := 0; i < numsLen; i++ {
		if nums[i] != -1 {
			cpNums = append(cpNums, nums[i])
		}
	}
	nums = cpNums
	return result
}
