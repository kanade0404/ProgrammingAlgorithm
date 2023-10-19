package squares_of_a_sorted_array

import (
	"math"
	"sort"
)

func sortedSquares(nums []int) []int {
	results := make([]int, len(nums), len(nums))
	for i := range nums {
		results[i] = int(math.Pow(float64(nums[i]), float64(2)))
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})
	return results
}
