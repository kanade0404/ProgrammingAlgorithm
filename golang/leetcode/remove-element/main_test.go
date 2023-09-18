package remove_element

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want:     2,
			wantNums: []int{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
				return
			}
			if diff := cmp.Diff(tt.args.nums, tt.wantNums); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
