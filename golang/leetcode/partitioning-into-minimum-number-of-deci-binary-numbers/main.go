package partitioning_into_minimum_number_of_deci_binary_numbers

import (
	"strconv"
	"strings"
)

func minPartitions(n string) int {
	arr := strings.Split(n, "")
	var result int
	for _, a := range arr {
		n, _ := strconv.Atoi(a)
		if n > result {
			result = n
		}
	}
	return result
}
