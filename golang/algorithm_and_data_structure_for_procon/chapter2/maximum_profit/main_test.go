package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		ns   []int
		want int
	}{
		{
			name: "test1",
			ns:   []int{6, 4, 3, 1, 3, 4, 3},
			want: 3,
		},
		{
			name: "test2",
			ns:   []int{3, 4, 3, 2},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.ns); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
