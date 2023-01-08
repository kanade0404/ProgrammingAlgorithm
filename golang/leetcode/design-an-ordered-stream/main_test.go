package design_an_ordered_stream

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestOrderedStream_Insert(t *testing.T) {
	type args struct {
		idKey int
		value string
	}
	type testCase struct {
		inserter args
		want     []string
	}
	tests := []struct {
		name      string
		n         int
		testCases []testCase
	}{
		{
			name: "test1",
			n:    5,
			testCases: []testCase{
				{
					inserter: args{
						3,
						"ccccc",
					},
					want: []string{},
				},
				{
					inserter: args{
						1,
						"aaaaa",
					},
					want: []string{"aaaaa"},
				},
				{
					inserter: args{
						2,
						"bbbbb",
					},
					want: []string{"bbbbb", "ccccc"},
				},
				{
					inserter: args{
						5,
						"eeeee",
					},
					want: []string{},
				},
				{
					inserter: args{
						4,
						"ddddd",
					},
					want: []string{"ddddd", "eeeee"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.n)
			for i, t2 := range tt.testCases {
				th := this.Insert(t2.inserter.idKey, t2.inserter.value)
				fmt.Printf("%+v\n", this)
				if diff := cmp.Diff(th, t2.want); diff != "" {
					t.Errorf("idx: %d\n(-got+want)\n%v", i, diff)
					return
				}
			}
		})
	}
}
