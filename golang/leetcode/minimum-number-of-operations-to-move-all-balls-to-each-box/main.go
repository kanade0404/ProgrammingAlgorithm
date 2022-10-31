package minimum_number_of_operations_to_move_all_balls_to_each_box

import (
	"math"
	"strings"
)

func minOperations(boxes string) []int {
	newBoxes := strings.Split(boxes, "")
	boxesLen := len(newBoxes)
	var moveCounts []int
	for i := 0; i < boxesLen; i++ {
		var moveCount int
		for j := 0; j < boxesLen; j++ {
			if i == j {
				continue
			}
			if newBoxes[j] == "1" {
				moveCount += int(math.Abs(float64(i - j)))
			}
		}
		moveCounts = append(moveCounts, moveCount)
	}
	return moveCounts
}
