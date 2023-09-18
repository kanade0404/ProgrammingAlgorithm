package min_stack

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMinStack_GetMin(t *testing.T) {
	tests := []struct {
		name string
		this *MinStack
		want int
	}{
		{
			name: "test1",
			this: &MinStack{
				arr: []int{1, 4, 2},
			},
			want:1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{}
			if got := this.GetMin(); got != tt.want {
				t.Errorf("GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMain(t *testing.T) {
	tests := []struct{
		name string
		operation string
		input int
		want int
	}{
		{},
	}
	this := &MinStack{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.operation {
			case "push":
				this.Push(tt.input)
				if diff := cmp.Diff(this, )
			case "min":
			case "pop":
			case "top":
			default:
				t.Errorf("this operation not found. %s", tt.operation)
			}
		})
	}
}
