package partitioning_into_minimum_number_of_deci_binary_numbers

import (
	"sort"
	"strconv"
	"strings"
)

func minPartitions(n string) int {
	arr := strings.Split(n, "")
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	result, _ := strconv.Atoi(arr[0])
	return result
}
