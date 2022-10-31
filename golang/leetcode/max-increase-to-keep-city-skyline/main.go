package max_increase_to_keep_city_skyline

import (
	"fmt"
)

func maxIncreaseKeepingSkyline(grid [][]int) int {
	var (
		rowMax, colMax []int
		gridReplica    [][]int
	)
	N := len(grid)
	cols := make([][]int, N)
	for _, row := range grid {
		var max int
		var replica []int
		for _, r := range row {
			if r > max {
				max = r
			}
			replica = append(replica, r)
		}
		rowMax = append(rowMax, max)
		for i := 0; i < N; i++ {
			cols[i] = append(cols[i], row[i])
		}
		gridReplica = append(gridReplica, replica)
	}
	for _, col := range cols {
		var max int
		for _, c := range col {
			if c > max {
				max = c
			}
		}
		colMax = append(colMax, max)
	}
	var result int
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if gridReplica[i][j] == rowMax[i] || gridReplica[i][j] == colMax[j] {
				continue
			}
			if rowMax[i] > colMax[j] {
				result += colMax[j] - gridReplica[i][j]
			} else {
				result += rowMax[i] - gridReplica[i][j]
			}
			fmt.Println(i, j, rowMax, rowMax[i], colMax, colMax[j], result)
		}
	}
	return result
}
