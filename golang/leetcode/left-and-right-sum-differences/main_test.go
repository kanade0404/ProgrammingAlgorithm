package left_and_right_sum_differences

import (
	"reflect"
	"testing"
)

func Test_leftRightDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{10, 4, 8, 3},
			},
			want: []int{15, 1, 11, 22},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftRightDifference(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leftRightDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
