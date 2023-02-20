package difference_between_element_sum_and_digit_sum_of_an_array

import "testing"

func Test_differenceOfSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 15, 6, 3},
			},
			want: 9,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				nums: []int{8, 9, 10},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfSum(tt.args.nums); got != tt.want {
				t.Errorf("differenceOfSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
