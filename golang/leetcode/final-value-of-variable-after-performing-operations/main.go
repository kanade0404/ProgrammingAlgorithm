package final_value_of_variable_after_performing_operations

func finalValueAfterOperations(operations []string) int {
	var plusCount, minusCount int
	for _, operation := range operations {
		if operation == "X++" || operation == "++X" {
			plusCount = plusCount + 1
		} else {
			minusCount = minusCount + 1
		}
	}
	return plusCount - minusCount
}
