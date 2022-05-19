package subrectangle_queries

import (
	"reflect"
	"testing"
)

/**
["SubrectangleQueries","getValue","updateSubrectangle","getValue","getValue","updateSubrectangle","getValue","getValue"]
[[[[1,2,1],[4,3,4],[3,2,1],[1,1,1]]],[0,2],[0,0,3,2,5],[0,2],[3,1],[3,0,3,2,10],[3,1],[0,2]]
*/
func TestConstructor(t *testing.T) {
	type args struct {
		rectangle [][]int
	}
	tests := []struct {
		name string
		args args
		want SubrectangleQueries
	}{
		{
			name: "test",
			args: args{
				rectangle: [][]int{
					{1, 2, 1},
					{4, 3, 4},
					{3, 2, 1},
					{1, 1, 1},
				},
			},
			want: SubrectangleQueries{
				Row: 4,
				Col: 3,
				Rectangle: [][]int{
					{1, 2, 1},
					{4, 3, 4},
					{3, 2, 1},
					{1, 1, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.rectangle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubrectangleQueries_GetValue(t *testing.T) {
	type fields struct {
		Row       int
		Col       int
		Rectangle [][]int
	}
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "test",
			fields: fields{
				Row: 4,
				Col: 3,
				Rectangle: [][]int{
					{1, 2, 1},
					{4, 3, 4},
					{3, 2, 1},
					{1, 1, 1},
				},
			},
			args: args{
				row: 0,
				col: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &SubrectangleQueries{
				Row:       tt.fields.Row,
				Col:       tt.fields.Col,
				Rectangle: tt.fields.Rectangle,
			}
			if got := this.GetValue(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubrectangleQueries_UpdateSubrectangle(t *testing.T) {
	type fields struct {
		Row       int
		Col       int
		Rectangle [][]int
	}
	type args struct {
		row1     int
		col1     int
		row2     int
		col2     int
		newValue int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SubrectangleQueries
	}{
		{
			name: "test",
			fields: fields{
				Row: 4,
				Col: 3,
				Rectangle: [][]int{
					{1, 2, 1},
					{4, 3, 4},
					{3, 2, 1},
					{1, 1, 1},
				},
			},
			args: args{
				row1:     0,
				col1:     0,
				row2:     3,
				col2:     2,
				newValue: 5,
			},
			want: &SubrectangleQueries{
				Row: 4,
				Col: 3,
				Rectangle: [][]int{
					{5, 5, 5},
					{5, 5, 5},
					{5, 5, 5},
					{5, 5, 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &SubrectangleQueries{
				Row:       tt.fields.Row,
				Col:       tt.fields.Col,
				Rectangle: tt.fields.Rectangle,
			}
			this.UpdateSubrectangle(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2, tt.args.newValue)
			if !reflect.DeepEqual(this, tt.want) {
				t.Errorf("UpdateSubrectangle() this: %v, want: %v", this, tt.want)
			}
		})
	}
}
