package running_sum_of_1d_array

func runningSum(nums []int) []int {
	var result []int
	var total int
	for i, num := range nums {
		if i == 0 {
			result = append(result, num)
		} else {
			result = append(result, total+num)
		}
		total = total + num
	}
	return result
}
