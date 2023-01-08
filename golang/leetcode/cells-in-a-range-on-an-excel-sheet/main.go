package cells_in_a_range_on_an_excel_sheet

import (
	"log"
	"strconv"
	"strings"
)

var (
	AtoZ    = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	AtoZMap = map[string]int{
		"A":   0,
		"B":   1,
		"C":   2,
		"D":   3,
		"E":   4,
		"F":   5,
		"G":   6,
		"H":   7,
		"I":   8,
		"J":   9,
		"K":   10,
		"L":   11,
		"M":   12,
		"N":   13,
		"O":   14,
		"P":   15,
		"Q":   16,
		"R":   17,
		"S":   18,
		"T":   19,
		"U":   20,
		"V":   21,
		"W":   22,
		"X":   23,
		"Y":   24,
		"Z":   25,
		"NON": 26,
	}
)

func cellsInRange(s string) []string {
	cells := strings.Split(s, ":")
	minCell := strings.Split(cells[0], "")
	minCol := minCell[0]
	minRow, err := strconv.Atoi(minCell[1])
	if err != nil {
		log.Fatalln(err)
	}
	maxCell := strings.Split(cells[1], "")
	maxCol := maxCell[0]
	maxRow, err := strconv.Atoi(maxCell[1])
	if err != nil {
		log.Fatalln(err)
	}
	columns := AtoZ[AtoZMap[minCol] : AtoZMap[maxCol]+1]
	var results []string
	for _, c := range columns {
		for i := minRow; i < maxRow+1; i++ {
			results = append(results, c+strconv.Itoa(i))
		}
	}
	return results
}
