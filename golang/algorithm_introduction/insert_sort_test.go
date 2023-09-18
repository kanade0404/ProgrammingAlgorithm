package algorithm_introduction

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestInsertSort_Sort(t *testing.T) {
	type fields struct {
		arr []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "test",
			fields: fields{
				arr: []int{5, 2, 4, 6, 1, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := &InsertSort{
				arr: tt.fields.arr,
			}
			is.Sort()
			if diff := cmp.Diff(tt.want, is.arr); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
