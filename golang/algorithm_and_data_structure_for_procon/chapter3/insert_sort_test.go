package chapter3

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{5, 2, 4, 6, 1, 3},
				n:    6,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := &Arr{tt.args.nums}
			arr.InsertSort(tt.args.n)
			if !reflect.DeepEqual(arr.Nums, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", arr.Nums, tt.want)
			}
		})
	}
}
