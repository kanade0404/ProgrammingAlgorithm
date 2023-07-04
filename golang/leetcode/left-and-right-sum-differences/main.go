package left_and_right_sum_differences

import (
	"math"
)

func leftRightDifference(nums []int) []int {
	var leftSum, rightSum, result []int
	for i := 0; i < len(nums); i++ {
		var left, right int
		if i == 0 {
			leftSum = append(leftSum, 0)
		} else {
			var lr int
			for li := range nums[:i] {
				lr += nums[:i][li]
			}
			left = lr
			leftSum = append(leftSum, left)
		}
		if i == len(nums)-1 {
			rightSum = append(rightSum, 0)
		} else {
			var rr int
			for ri := range nums[i+1:] {
				rr += nums[i+1:][ri]
			}
			right = rr
			rightSum = append(rightSum, right)
		}
		result = append(result, int(math.Abs(float64(left-right))))
	}
	return result
}
