package difference_between_element_sum_and_digit_sum_of_an_array

import "strconv"

func differenceOfSum(nums []int) int {
	var elementSum, digitSum int
	for i := range nums {
		elementSum += nums[i]
		if nums[i] > 9 {
			strN := strconv.Itoa(nums[i])
			for i := range strN {
				n, _ := strconv.Atoi(string(strN[i]))
				digitSum += n
			}
		} else {
			digitSum += nums[i]
		}
	}
	return elementSum - digitSum
}
