package number_of_steps_to_reduce_a_number_to_zero

func numberOfSteps(num int) int {
	var count int
	result := num
	for result > 0 {
		if result%2 == 0 {
			result /= 2
			count++
		}
		if result%2 == 1 {
			result -= 1
			count++
		}
	}
	return count
}
