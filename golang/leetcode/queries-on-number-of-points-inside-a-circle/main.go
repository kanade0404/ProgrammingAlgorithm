package queries_on_number_of_points_inside_a_circle

import "math"

func countPoints(points [][]int, queries [][]int) []int {
	var results []int
	for _, query := range queries {
		var count int
		for _, point := range points {
			if math.Sqrt(math.Pow(float64(query[0]-point[0]), 2)+math.Pow(float64(query[1]-point[1]), 2)) <= float64(query[2]) {
				count += 1
			}
		}
		results = append(results, count)
	}
	return results
}
