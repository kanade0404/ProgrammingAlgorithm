package partitioning_into_minimum_number_of_deci_binary_numbers

import "testing"

func Test_minPartitions(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				n: "32",
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				n: "82734",
			},
			want: 8,
		},
		{
			name: "test3",
			args: args{
				n: "27346209830709182346",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPartitions(tt.args.n); got != tt.want {
				t.Errorf("minPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
