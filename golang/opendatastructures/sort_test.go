package opendatastructures

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestArr_InsertSort(t *testing.T) {
	tests := []struct {
		name     string
		a        Arr
		expected Arr
	}{
		{
			name:     "test",
			a:        []int{5, 2, 4, 6, 1, 3},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.InsertSort()
			if diff := cmp.Diff(tt.a, tt.expected); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
