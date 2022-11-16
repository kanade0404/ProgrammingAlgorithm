package how_many_numbers_are_smaller_than_the_current_number

func smallerNumbersThanCurrent(nums []int) []int {
	var results []int
	for i := range nums {
		var count int
		for j := range nums {
			if nums[i] > nums[j] {
				count += 1
			}
		}
		results = append(results, count)
	}
	return results
}
