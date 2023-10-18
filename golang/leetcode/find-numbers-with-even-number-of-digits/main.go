package find_numbers_with_even_number_of_digits

import "strconv"

func findNumbers(nums []int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		if len(strconv.Itoa(nums[i]))%2 == 0 {
			count++
		}
	}
	return count
}
