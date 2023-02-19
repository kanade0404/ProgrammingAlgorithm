package contains_duplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = struct{}{}
		}
	}
	return false
}
