package subrectangle_queries

type SubrectangleQueries struct {
	Row       int
	Col       int
	Rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	row := len(rectangle)
	col := len(rectangle[0])
	return SubrectangleQueries{
		Row:       row,
		Col:       col,
		Rectangle: rectangle,
	}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i < row2+1; i++ {
		for j := col1; j < col2+1; j++ {
			this.Rectangle[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.Rectangle[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
