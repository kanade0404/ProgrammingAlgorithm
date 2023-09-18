package create_target_array_in_the_given_order

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	l := len(nums)
	results := make([]int, l)
	for i := 0; i < l; i++ {
		n := nums[i]
		idx := index[i]
		b := results[:idx]
		if results[idx] == n {
			results[idx] = n
		} else {
			before := make([]int, len(b))
			copy(before, b)
			a := results[idx:]
			after := make([]int, len(a))
			copy(after, a)
			r := make([]int, len(b))
			r = append(r, before...)
			r = append(r, n)
			r = append(r, after...)
			results = r
			fmt.Println(results, n, idx)
		}
	}
	return results
}
