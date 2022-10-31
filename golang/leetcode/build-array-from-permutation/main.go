package build_array_from_permutation

func buildArray(nums []int) []int {
	var result []int
	for _, num := range nums {
		result = append(result, nums[num])
	}
	return result
}
