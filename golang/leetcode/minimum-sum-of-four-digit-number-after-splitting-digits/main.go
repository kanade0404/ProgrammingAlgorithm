package minimum_sum_of_four_digit_number_after_splitting_digits

import (
	"sort"
	"strconv"
	"strings"
)

func minimumSum(num int) int {
	arr := strings.Split(strconv.Itoa(num), "")
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	a1, _ := strconv.Atoi(arr[0] + arr[2])
	a2, _ := strconv.Atoi(arr[1] + arr[3])
	return a1 + a2
}
