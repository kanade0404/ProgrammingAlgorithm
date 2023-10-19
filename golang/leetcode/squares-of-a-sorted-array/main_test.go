package squares_of_a_sorted_array

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{
				[]int{-4, -1, 0, 3, 10},
			},
			[]int{0, 1, 9, 16, 100},
		},
		{
			"test2",
			args{
				[]int{-7, -3, 2, 3, 11},
			},
			[]int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
