package binary_search

func search(nums []int, target int) int {
	l := len(nums)
	if nums[l-1] < target {
		return -1
	}
	low := 0
	high := l - 1

	for low <= high {
		median := (low + high) / 2
		if nums[median] == target {
			return median
		} else if nums[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return -1
}
