package cells_in_a_range_on_an_excel_sheet

import (
	"reflect"
	"testing"
)

func Test_cellsInRange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				s: "K1:L2",
			},
			want: []string{"K1", "K2", "L1", "L2"},
		},
		{
			name: "test2",
			args: args{
				s: "A1:F1",
			},
			want: []string{"A1", "B1", "C1", "D1", "E1", "F1"},
		},
		{
			name: "test3",
			args: args{
				s: "U7:X9",
			},
			want: []string{"U7", "U8", "U9", "V7", "V8", "V9", "W7", "W8", "W9", "X7", "X8", "X9"},
		},
		{
			name: "test4",
			args: args{
				s: "A1:Z1",
			},
			want: []string{"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1", "I1", "J1", "K1", "L1", "M1", "N1", "O1", "P1", "Q1", "R1", "S1", "T1", "U1", "V1", "W1", "X1", "Y1", "Z1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cellsInRange(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cellsInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
