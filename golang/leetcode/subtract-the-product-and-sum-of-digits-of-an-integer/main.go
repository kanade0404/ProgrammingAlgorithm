package subtract_the_product_and_sum_of_digits_of_an_integer

import (
	"strconv"
	"strings"
)

func subtractProductAndSum(n int) int {
	arr := strings.Split(strconv.Itoa(n), "")
	var productResult, sumResult = 1, 0
	for i := range arr {
		num, _ := strconv.Atoi(arr[i])
		productResult = productResult * num
		sumResult = sumResult + num

	}
	return productResult - sumResult
}
